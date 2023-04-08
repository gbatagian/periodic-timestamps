package ptlist

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"periodic-timestamps/settings"
	"periodic-timestamps/utils"
	"testing"

	"github.com/gin-gonic/gin"
)

var baseURL string
var testRouter *gin.Engine

func init() {
	baseURL = "/ptlist"
	testRouter = settings.ENGINE
	testRouter.GET(baseURL, PtListGet())
}

func formatURLParameters(inputs map[string]string) string {
	values := url.Values{}
	for k, v := range inputs {
		if v != "" {
			values.Add(k, v)
		}
	}
	return "?" + values.Encode()
}

func TestOKHourly(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1h",
		"tz":     "Europe/Athens",
		"t1":     "20210714T204603Z",
		"t2":     "20210715T123456Z",
	}
	URL := baseURL + formatURLParameters(paramValues)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	// Act
	testRouter.ServeHTTP(resp, req)

	// Assert
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Request failed with status code: %d", status)
	}
	if body := resp.Body.String(); body != utils.ToJsonString(
		[]string{
			"20210714T210000Z",
			"20210714T220000Z",
			"20210714T230000Z",
			"20210715T000000Z",
			"20210715T010000Z",
			"20210715T020000Z",
			"20210715T030000Z",
			"20210715T040000Z",
			"20210715T050000Z",
			"20210715T060000Z",
			"20210715T070000Z",
			"20210715T080000Z",
			"20210715T090000Z",
			"20210715T100000Z",
			"20210715T110000Z",
			"20210715T120000Z",
		},
	) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}
}

func TestOKDaily(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1d",
		"tz":     "Europe/Athens",
		"t1":     "20211010T204603Z",
		"t2":     "20211115T123456Z",
	}
	URL := baseURL + formatURLParameters(paramValues)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	// Act
	testRouter.ServeHTTP(resp, req)

	// Assert
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Request failed with status code: %d", status)
	}
	if body := resp.Body.String(); body != utils.ToJsonString(
		[]string{
			"20211010T210000Z",
			"20211011T210000Z",
			"20211012T210000Z",
			"20211013T210000Z",
			"20211014T210000Z",
			"20211015T210000Z",
			"20211016T210000Z",
			"20211017T210000Z",
			"20211018T210000Z",
			"20211019T210000Z",
			"20211020T210000Z",
			"20211021T210000Z",
			"20211022T210000Z",
			"20211023T210000Z",
			"20211024T210000Z",
			"20211025T210000Z",
			"20211026T210000Z",
			"20211027T210000Z",
			"20211028T210000Z",
			"20211029T210000Z",
			"20211030T210000Z",
			"20211031T220000Z",
			"20211101T220000Z",
			"20211102T220000Z",
			"20211103T220000Z",
			"20211104T220000Z",
			"20211105T220000Z",
			"20211106T220000Z",
			"20211107T220000Z",
			"20211108T220000Z",
			"20211109T220000Z",
			"20211110T220000Z",
			"20211111T220000Z",
			"20211112T220000Z",
			"20211113T220000Z",
			"20211114T220000Z",
		},
	) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}
}

func TestOKMonthly(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1mo",
		"tz":     "Europe/Athens",
		"t1":     "20210214T204603Z",
		"t2":     "20211115T123456Z",
	}
	URL := baseURL + formatURLParameters(paramValues)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	// Act
	testRouter.ServeHTTP(resp, req)

	// Assert
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Request failed with status code: %d", status)
	}
	if body := resp.Body.String(); body != utils.ToJsonString(
		[]string{
			"20210228T220000Z",
			"20210331T210000Z",
			"20210430T210000Z",
			"20210531T210000Z",
			"20210630T210000Z",
			"20210731T210000Z",
			"20210831T210000Z",
			"20210930T210000Z",
			"20211031T220000Z",
		},
	) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}
}

func TestOKYearly(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1y",
		"tz":     "Europe/Athens",
		"t1":     "20180214T204603Z",
		"t2":     "20211115T123456Z",
	}
	URL := baseURL + formatURLParameters(paramValues)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	// Act
	testRouter.ServeHTTP(resp, req)

	// Assert
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Request failed with status code: %d", status)
	}
	if body := resp.Body.String(); body != utils.ToJsonString(
		[]string{
			"20181231T220000Z",
			"20191231T220000Z",
			"20201231T220000Z",
		},
	) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}
}

func Test400InvalidPeriodValue(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1w",
		"tz":     "Europe/Athens",
		"t1":     "20180214T204603Z",
		"t2":     "20211115T123456Z",
	}
	URL := baseURL + formatURLParameters(paramValues)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	// Act
	testRouter.ServeHTTP(resp, req)

	// Assert
	if status := resp.Code; status != http.StatusBadRequest {
		t.Errorf("Response status code is not 400. Status code: %d", status)
	}
	if body := resp.Body.String(); body != utils.ToJsonString(map[string]interface{}{
		"status": "error",
		"desc":   "Unsupported period",
	}) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}

}

func Test400ParameterMissing(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1h",
		"tz":     "Europe/Athens",
		"t1":     "20210714T204603Z",
		"t2":     "20210715T123456Z",
	}

	for k := range paramValues {
		paramValueBackUp := paramValues[k]
		paramValues[k] = ""

		URL := baseURL + formatURLParameters(paramValues)
		req, err := http.NewRequest("GET", URL, nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		// Act
		testRouter.ServeHTTP(resp, req)

		// Assert
		if status := resp.Code; status != http.StatusBadRequest {
			t.Errorf("Response status code is not 400. Status code: %d", status)
		}
		if body := resp.Body.String(); body != utils.ToJsonString(map[string]interface{}{
			"status": "error",
			"desc":   fmt.Sprintf("The '%s' parameter cannot be empty. Please provide a value for this parameter.", k),
		}) {
			t.Errorf("Response payload has invalid format. Body: %s", body)
		}

		paramValues[k] = paramValueBackUp
	}
}

func Test400InvalidDateTimeFormat(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1y",
		"tz":     "Europe/Athens",
		"t1":     "20180214T204603Z",
		"t2":     "20210715T123456",
	}
	URL := baseURL + formatURLParameters(paramValues)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	// Act
	testRouter.ServeHTTP(resp, req)

	// Assert
	if status := resp.Code; status != http.StatusBadRequest {
		t.Errorf("Response status code is not 400")
	}
	if body := resp.Body.String(); body != utils.ToJsonString(map[string]interface{}{
		"status": "error",
		"desc":   fmt.Sprintf("Timestamp '%s' does not match the required format YYYYDDMMTHHMMSSZ", paramValues["t2"]),
	}) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}

}

func Test400InvalidTimezone(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1y",
		"tz":     "foo",
		"t1":     "20180214T204603Z",
		"t2":     "20210715T123456Z",
	}
	URL := baseURL + formatURLParameters(paramValues)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()

	// Act
	testRouter.ServeHTTP(resp, req)

	// Assert
	if status := resp.Code; status != http.StatusBadRequest {
		t.Errorf("Response status code is not 400")
	}
	if body := resp.Body.String(); body != utils.ToJsonString(map[string]interface{}{
		"status": "error",
		"desc":   fmt.Sprintf("unknown time zone %s", paramValues["tz"]),
	}) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}

}
