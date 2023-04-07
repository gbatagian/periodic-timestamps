package ptlist

import (
	"testing"
	"time"
)

func TestPeriodicTimestampsYearly(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		"1y",
		"20230311T111111Z", // 2023-03-11 11:11:11 UTC
		"20240311T111111Z", // 2024-03-11 11:11:11 UTC
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
		time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		"1m",
		"20230311T111111Z", // 2023-03-11 11:11:11 UTC
		"20231201T121212Z", // 2023-12-01 12:12:12 UTC
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

func TestPeriodicTimestampsDay(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 12, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 13, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 14, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 16, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 17, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 18, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 19, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 22, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 24, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 25, 0, 0, 0, 0, time.UTC),
	}

	// Act
	ptList := PeriodicTimestamps(
		"1d",
		"20230311T111111Z", // 2023-03-11 11:11:11 UTC
		"20230325T121212Z", // 2023-03-25 12:12:12 UTC
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

func TestPeriodicTimestampsHour(t *testing.T) {
	// Arrange
	expectedPtList := []time.Time{
		time.Date(2023, 3, 11, 11, 0, 0, 0, time.UTC),
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
		"1h",
		"20230311T111111Z", // 2023-03-11 11:11:11 UTC
		"20230311T211212Z", // 2023-03-25 21:12:12 UTC
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
