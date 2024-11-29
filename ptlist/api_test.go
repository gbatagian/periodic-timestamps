package ptlist

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"periodic-timestamps/utils"
	"reflect"
	"strings"
	"testing"
)

var baseURL string
var testRouter *http.ServeMux
var okTestCases []struct {
	name             string
	paramValues      map[string]string
	expectedRespJSON []string
}
var badRequestTestCases []struct {
	name             string
	paramValues      map[string]string
	expectedRespBody string
}

func init() {
	baseURL = "/ptlist"
	testRouter = http.NewServeMux()
	testRouter.HandleFunc("GET /ptlist", PtListGet())

	okTestCases = []struct {
		name             string
		paramValues      map[string]string
		expectedRespJSON []string
	}{
		{
			name: "Hourly OK",
			paramValues: map[string]string{
				"period": "1h",
				"tz":     "Europe/Athens",
				"t1":     "20210714T204603Z",
				"t2":     "20210715T123456Z",
			},
			expectedRespJSON: []string{
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
		},
		{
			name: "Daily OK",
			paramValues: map[string]string{
				"period": "1d",
				"tz":     "Europe/Athens",
				"t1":     "20211010T204603Z",
				"t2":     "20211115T123456Z",
			},
			expectedRespJSON: []string{
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
		},
		{
			name: "Monthly OK",
			paramValues: map[string]string{
				"period": "1mo",
				"tz":     "Europe/Athens",
				"t1":     "20210214T204603Z",
				"t2":     "20211115T123456Z",
			},
			expectedRespJSON: []string{
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
		},
		{
			name: "Yearly OK",
			paramValues: map[string]string{
				"period": "1y",
				"tz":     "Europe/Athens",
				"t1":     "20180214T204603Z",
				"t2":     "20211115T123456Z",
			},
			expectedRespJSON: []string{
				"20181231T220000Z",
				"20191231T220000Z",
				"20201231T220000Z",
			},
		},
	}

	badRequestTestCases = []struct {
		name             string
		paramValues      map[string]string
		expectedRespBody string
	}{
		{
			name: "Invalid period value",
			paramValues: map[string]string{
				"period": "1w",
				"tz":     "Europe/Athens",
				"t1":     "20180214T204603Z",
				"t2":     "20211115T123456Z",
			},
			expectedRespBody: "unsupported period '1w'",
		},
		{
			name: "Invalid datetime format",
			paramValues: map[string]string{
				"period": "1y",
				"tz":     "Europe/Athens",
				"t1":     "20180214T204603Z",
				"t2":     "20210715T123456",
			},
			expectedRespBody: "timestamp '20210715T123456' does not match the required format YYYYDDMMTHHMMSSZ",
		},
		{
			name: "Invalid timezone format",
			paramValues: map[string]string{
				"period": "1y",
				"tz":     "foo",
				"t1":     "20180214T204603Z",
				"t2":     "20210715T123456Z",
			},
			expectedRespBody: "unknown time zone foo",
		},
		{
			name: "Parameter period missing",
			paramValues: map[string]string{
				"tz": "Europe/Athens",
				"t1": "20210714T204603Z",
				"t2": "20210715T123456Z",
			},
			expectedRespBody: "parameter 'period' cannot be empty, please provide a value",
		},
		{
			name: "Parameter tz missing",
			paramValues: map[string]string{
				"period": "1d",
				"t1":     "20210714T204603Z",
				"t2":     "20210715T123456Z",
			},
			expectedRespBody: "parameter 'tz' cannot be empty, please provide a value",
		},
		{
			name: "Parameter t1 missing",
			paramValues: map[string]string{
				"period": "1d",
				"tz":     "Europe/Athens",
				"t2":     "20210715T123456Z",
			},
			expectedRespBody: "parameter 't1' cannot be empty, please provide a value",
		},
		{
			name: "Parameter t2 missing",
			paramValues: map[string]string{
				"period": "1d",
				"tz":     "Europe/Athens",
				"t1":     "20210714T204603Z",
			},
			expectedRespBody: "parameter 't2' cannot be empty, please provide a value",
		},
	}
}

func TestOK(t *testing.T) {
	for _, testCase := range okTestCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				// Arrange
				paramValues := testCase.paramValues
				url := baseURL + utils.FormatURLParameters(paramValues)
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					t.Fatal(err)
				}
				resp := httptest.NewRecorder()

				// Act
				testRouter.ServeHTTP(resp, req)
				var responseJSON []string
				json.NewDecoder(resp.Body).Decode(&responseJSON)

				// Assert
				if status := resp.Code; status != http.StatusOK {
					t.Errorf("Request failed with status code: %d", status)
				}
				if !reflect.DeepEqual(responseJSON, testCase.expectedRespJSON) {
					t.Errorf("Response payload has invalid format. Body: %v", responseJSON)
				}
			},
		)
	}
}

func TestBadRequest(t *testing.T) {
	for _, testCase := range badRequestTestCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				// Arrange
				paramValues := testCase.paramValues
				url := baseURL + utils.FormatURLParameters(paramValues)
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					t.Fatal(err)
				}
				resp := httptest.NewRecorder()

				// Act
				testRouter.ServeHTTP(resp, req)
				respBody, _ := io.ReadAll(resp.Body)

				// Assert
				if status := resp.Code; status != http.StatusBadRequest {
					t.Errorf("Request failed with status code: %d", status)
				}
				if !(strings.TrimSpace(string(respBody)) == testCase.expectedRespBody) {
					t.Errorf("Invalid response body: %v", respBody)
				}
			},
		)
	}
}
