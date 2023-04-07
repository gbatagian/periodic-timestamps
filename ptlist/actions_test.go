package ptlist

import (
	"testing"
	"time"
)

func TestPeriodicTimestampsYearly(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 12, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2024, 12, 31, 12, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		Yearly,
		time.Date(2023, 3, 11, 11, 11, 11, 0, time.UTC),
		time.Date(2025, 3, 11, 11, 11, 11, 0, time.UTC),
		time.UTC,
	)

	// Assert
	if len(ptList) != len(expectedPtList) {
		t.Errorf("Invalid periods list. Expected: %v", ptList)
	}

	for idx, e := range ptList {
		if expectedPtList[idx] != e {
			t.Errorf("Invalid periods list. Expected: %v", ptList)
		}
	}
}

func TestPeriodicTimestampsMonthly(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 1, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 2, 28, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 30, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 5, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 6, 30, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 7, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 8, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 9, 30, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 11, 30, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 31, 12, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		Monthly,
		time.Date(2023, 1, 11, 11, 11, 11, 0, time.UTC),
		time.Date(2024, 2, 5, 15, 15, 15, 0, time.UTC),
		time.UTC,
	)

	// Assert
	if len(ptList) != len(expectedPtList) {
		t.Errorf("Invalid periods list. Expected: %v", ptList)
	}

	for idx, e := range ptList {
		if expectedPtList[idx] != e {
			t.Errorf("Invalid periods list. Expected: %v", ptList)
		}
	}
}

func TestPeriodicTimestampsDaily(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 3, 25, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 26, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 27, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 28, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 29, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 30, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 31, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 2, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 3, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 4, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 5, 12, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		Daily,
		time.Date(2023, 3, 25, 11, 11, 11, 0, time.UTC),
		time.Date(2023, 4, 5, 15, 15, 15, 0, time.UTC),
		time.UTC,
	)

	// Assert
	if len(ptList) != len(expectedPtList) {
		t.Errorf("Invalid periods list. Expected: %v", ptList)
	}

	for idx, e := range ptList {
		if expectedPtList[idx] != e {
			t.Errorf("Invalid periods list. Expected: %v", ptList)
		}
	}
}

func TestPeriodicTimestampsHourly(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 3, 11, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 13, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 14, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 15, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 16, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 17, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 18, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 19, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 20, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 21, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		Hourly,
		time.Date(2023, 3, 11, 11, 11, 11, 0, time.UTC),
		time.Date(2023, 3, 11, 21, 21, 21, 0, time.UTC),
		time.UTC,
	)

	// Assert
	if len(ptList) != len(expectedPtList) {
		t.Errorf("Invalid periods list. Expected: %v", ptList)
	}

	for idx, e := range ptList {
		if expectedPtList[idx] != e {
			t.Errorf("Invalid periods list. Expected: %v", ptList)
		}
	}
}

// Assert that when the start of the period is in an exact time
// that time is not included in the list of timestamps
func TestPeriodicTimestampsHourlyStartOnWholeTime(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 3, 11, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 13, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 14, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 15, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 16, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 17, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 18, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 19, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 20, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 21, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		Hourly,
		time.Date(2023, 3, 11, 11, 00, 00, 0, time.UTC),
		time.Date(2023, 3, 11, 21, 21, 21, 0, time.UTC),
		time.UTC,
	)

	// Assert
	if len(ptList) != len(expectedPtList) {
		t.Errorf("Invalid periods list. Expected: %v", ptList)
	}

	for idx, e := range ptList {
		if expectedPtList[idx] != e {
			t.Errorf("Invalid periods list. Expected: %v", ptList)
		}
	}
}

// Assert that when the end of the period is in an exact time
// that time is not included in the list of timestamps
func TestPeriodicTimestampsHourlyEndOnWholeTime(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 3, 11, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 13, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 14, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 15, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 16, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 17, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 18, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 19, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 11, 20, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		Hourly,
		time.Date(2023, 3, 11, 11, 00, 00, 0, time.UTC),
		time.Date(2023, 3, 11, 21, 00, 00, 0, time.UTC),
		time.UTC,
	)

	// Assert
	if len(ptList) != len(expectedPtList) {
		t.Errorf("Invalid periods list. Expected: %v", ptList)
	}

	for idx, e := range ptList {
		if expectedPtList[idx] != e {
			t.Errorf("Invalid periods list. Expected: %v", ptList)
		}
	}
}
