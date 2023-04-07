package ptlist

import (
	"time"
)

func PeriodicTimestamps(period string, timeString1 string, timeString2 string) []time.Time {
	p, _ := PeriodFromString(period)
	layout := "20060102T150405Z"

	t1, _ := time.Parse(layout, timeString1)
	sopTime1 := toStartOfPeriod(p, t1)

	t2, _ := time.Parse(layout, timeString2)
	sopTime2 := toStartOfPeriod(p, t2)

	return periodicTimestamps(p, sopTime1, sopTime2)
}

func toStartOfPeriod(p Period, t time.Time) time.Time {
	year, month, day, hour := t.Year(), t.Month(), t.Day(), t.Hour()

	switch p {
	case Yearly:
		return time.Date(year, time.January, 1, 0, 0, 0, 0, t.Location())

	case Monthly:
		return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())

	case Daily:
		return time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	default:
		return time.Date(year, month, day, hour, 0, 0, 0, t.Location())
	}
}

func periodicTimestamps(p Period, t1 time.Time, t2 time.Time) []time.Time {
	var pt []time.Time

	t := t1
	pt = append(pt, t)

	switch p {
	case Yearly:
		for t.Before(t2) {
			t = t.AddDate(1, 0, 0)
			pt = append(pt, t)
		}
		return pt
	case Monthly:
		for t.Before(t2) {
			t = t.AddDate(0, 1, 0)
			pt = append(pt, t)
		}
		return pt
	case Daily:
		for t.Before(t2) {
			t = t.AddDate(0, 0, 1)
			pt = append(pt, t)
		}
		return pt
	default:
		for t.Before(t2) {
			t = t.Add(time.Hour)
			pt = append(pt, t)
		}
		return pt
	}
}
