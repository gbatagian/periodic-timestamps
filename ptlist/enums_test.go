package ptlist

import (
	"testing"
)

func TestStringFromPeriodOK(t *testing.T) {
	testCases := []struct {
		period   Period
		expected string
	}{
		{Hourly, "1h"},
		{Daily, "1d"},
		{Monthly, "1mo"},
		{Yearly, "1y"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			result := tc.period.String()
			if result != tc.expected {
				t.Errorf("String() for %v: expected %q, got %q", tc.period, tc.expected, result)
			}
		})
	}
}

func TestPeriodStringInvalidValue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The function did not panic as expected")
		}
	}()
	_ = Period.String(5)
}

func TestPeriodFromStringOK(t *testing.T) {
	testCases := []struct {
		inStr    string
		expected Period
	}{
		{"1h", Hourly},
		{"1d", Daily},
		{"1mo", Monthly},
		{"1y", Yearly},
	}

	for _, tc := range testCases {
		t.Run(tc.inStr, func(t *testing.T) {
			result, err := PeriodFromString(tc.inStr)
			if err != nil {
				t.Errorf("Unsupported period: %s", result)
			}
			if result != tc.expected {
				t.Errorf("Unexpected value %d for string %s", result, tc.inStr)
			}
		})
	}
}

func TestPeriodFromStringInvalidValue(t *testing.T) {
	p, err := PeriodFromString("aaa")
	if err == nil {
		t.Errorf("unsupportedPeriodError not raise")
	}
	if p != UnknownPeriod {
		t.Errorf("Invalid value returned for unknown period")
	}
}
