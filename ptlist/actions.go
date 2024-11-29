package ptlist

import (
	"time"
)

// PeriodicTimestamps generates a slice of timestamps between two given timestamps,
// in intervals based on the specified period. The timezone can also be specified.
// p: Period (Hourly, Daily, Monthly, Yearly)
// t1: Start timestamp
// t2: End timestamp
// tz: Timezone to apply to timestamps
func PeriodicTimestamps(p Period, t1 time.Time, t2 time.Time, tz *time.Location) []time.Time {
	invTms := invocationTimestamp(p, t1) // invocation timestamp based on the period
	return periodicTimestampsGenerator(p, invTms, t2, tz)
}

// invocationTimestamp calculates the start timestamp for a given period based
// on the input timestamp
func invocationTimestamp(p Period, t time.Time) time.Time {
	year, month, day, hour, minute := t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()

	toMonthlyInvocationTimestamp := func(t time.Time) time.Time {
		// Monthly invocation timestamp is the last day of month of the given timestamp.
		// -> Move to the first day of the month, add one month and then subtract one day
		//    to obtain the last day of the original month.
		//    For example, 2021-02-14 would become 2021-02-28.
		t = time.Date(t.Year(), t.Month(), 1, t.Hour(), 0, 0, 0, t.Location())
		t = t.AddDate(0, 1, -1)
		return t
	}

	invocationTime := map[Period]time.Time{
		Yearly:  time.Date(year, time.December, 31, hour, 0, 0, 0, t.Location()),
		Monthly: toMonthlyInvocationTimestamp(t),
		Daily:   time.Date(year, month, day, hour, 0, 0, 0, t.Location()),
		Hourly:  time.Date(year, month, day, hour, 0, 0, 0, t.Location()),
	}

	if sp, ok := invocationTime[p]; ok {
		if minute > 0 || p == Hourly {
			sp = sp.Add(time.Hour)
		}
		return sp
	}

	return time.Time{}
}

// periodicTimestampsGenerator creates a slice of timestamps based on the specified period
// between two input timestamps, in the provided timezone
func periodicTimestampsGenerator(p Period, t1 time.Time, t2 time.Time, tz *time.Location) []time.Time {
	var pt []time.Time
	t := t1

	DSTAdjustedTime := func(t time.Time) time.Time {
		if !t.In(tz).IsDST() && tz != time.UTC {
			return t.Add(time.Hour)
		}
		return t
	}

	switch p {
	case Yearly:
		for t.Before(t2) {
			pt = append(pt, DSTAdjustedTime(t))
			t = t.AddDate(1, 0, 0)
		}
		return pt
	case Monthly:
		for t.Before(t2) {
			pt = append(pt, DSTAdjustedTime(t))

			// Next element to append is the last day of the next month
			// --> Move to the first day of the current month, add two months and then subtract one day
			//     to obtain the last day of the original month's next month.
			//     For example, 2021-14-28 would become 2021-03-31.
			t = time.Date(t.Year(), t.Month(), 1, t.Hour(), 0, 0, 0, t.Location())
			t = t.AddDate(0, 2, -1)
		}
		return pt
	case Daily:
		for t.Before(t2) {
			pt = append(pt, DSTAdjustedTime(t))
			t = t.AddDate(0, 0, 1)
		}
		return pt
	default:
		for t.Before(t2) {
			pt = append(pt, DSTAdjustedTime(t))
			t = t.Add(time.Hour)
		}
		return pt
	}
}
