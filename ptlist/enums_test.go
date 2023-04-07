package ptlist

import (
	"testing"
)

func TestPeriodStringOK(t *testing.T) {
	values := []Period{
		Hourly,
		Daily,
		Monthly,
		Yearly,
	}
	for _, v := range values {
		Period.String(v)
	}
}

func TestPeriodStringInvalidValue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The function did not panic as expected")
		}
	}()

	Period.String(5)
}

func TestPeriodFromStringOK(t *testing.T) {

	expectedValues := []Period{
		Hourly,
		Daily,
		Monthly,
		Yearly,
	}
	for idx, ps := range []string{
		"1h", "1d", "1mo", "1y",
	} {
		v, err := PeriodFromString(ps)
		if err != nil {
			t.Errorf("Unsupported period: %s", v)
		}
		if v != expectedValues[idx] {
			t.Errorf("Unexpected value %d for string %s", v, ps)
		}
	}

}

func TestPeriodFromStringInvalidValue(t *testing.T) {
	_, err := PeriodFromString("aaa")
	if err == nil {
		t.Errorf("unsupportedPeriodError not raise")
	}
}
