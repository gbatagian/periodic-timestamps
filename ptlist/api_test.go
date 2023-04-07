package ptlist

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

func Test400ParameterMissing(t *testing.T) {
	// Arrange
	paramValues := map[string]string{
		"period": "1h",
		"tz":     "Europe/Athens",
		"t1":     "20210714T204603Z",
		"t2":     "20210715T123456Z",
	}
	formatURLParameters := func(inputs map[string]string) string {
		s := "?"
		for k, v := range inputs {
			if v != "" {
				if s == "?" {
					s = s + fmt.Sprintf("%s=%s", k, v)
				} else {
					s = s + "&" + fmt.Sprintf("%s=%s", k, v)
				}
			}
		}
		return s
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

func Test400InvalidPeriodValue(t *testing.T) {
	// Arrange
	invalidPeriodValue := "11h"
	URL := baseURL + fmt.Sprintf("?period=%s&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z", invalidPeriodValue)
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
		"desc":   "Unsupported period",
		"status": "error",
	}) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}

}

func Test400InvalidDateTimeFormat(t *testing.T) {
	// Arrange
	invalidDateTime := "20210715T123456"
	URL := baseURL + fmt.Sprintf("?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=%s", invalidDateTime)
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
		"desc":   fmt.Sprintf("Timestamp '%s' does not match the required format YYYYDDMMTHHMMSSZ", invalidDateTime),
	}) {
		t.Errorf("Response payload has invalid format. Body: %s", body)
	}

}
