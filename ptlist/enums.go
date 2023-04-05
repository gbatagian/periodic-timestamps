package ptlist

import "fmt"

type Period int

const (
	Hourly Period = iota
	Daily
	Weekly
	Monthly
	Yearly
)

func (p Period) String() string {
	names := []string{
		"1h", "1d", "1w", "1m", "1y",
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

func PeriodFromString(s string) Period {
	periodsMap := map[string]Period{
		"1h": Hourly,
		"1d": Daily,
		"1w": Weekly,
		"1m": Monthly,
		"1y": Yearly,
	}
	period, ok := periodsMap[s]
	if !ok {
		panic(
			fmt.Sprintf(
				"Invalid value '%s' provided as period. Please provide one of the supported values: [1h, 1w, 1d, 1m, 1y]", s,
			),
		)
	}
	return period
}
