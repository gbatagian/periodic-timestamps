package ptlist

import "fmt"

type Period int

const (
	Hourly Period = iota
	Daily
	Monthly
	Yearly
	UnknownPeriod
)

func (p Period) String() string {
	names := []string{
		"1h", "1d", "1mo", "1y",
	}
	if p < Hourly || p > Yearly {
		panic(
			fmt.Sprintf(
				"Invalid value %d provided as period. Periods should be in range [%d, %d]", p, Hourly, Yearly,
			),
		)
	}
	return names[p]
}

func PeriodFromString(s string) (Period, error) {
	periodsMap := map[string]Period{
		"1h":  Hourly,
		"1d":  Daily,
		"1mo": Monthly,
		"1y":  Yearly,
	}
	period, ok := periodsMap[s]
	if !ok {
		return UnknownPeriod, unsupportedPeriodError{}
	}
	return period, nil
}
