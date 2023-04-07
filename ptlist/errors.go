package ptlist

import "fmt"

type unsupportedPeriodError struct{}

func (e unsupportedPeriodError) Error() string {
	return "Unsupported period"
}

type invalidTimestampFormat struct {
	timestamp string
	format    string
}

func (e invalidTimestampFormat) Error() string {
	if e.format == "" {
		e.format = "YYYYDDMMTHHMMSSZ"
	}
	return fmt.Sprintf(
		"Timestamp '%s' does not match the required format %s", e.timestamp, e.format,
	)
}

type emptyQueryParameter struct {
	name string
}

func (e emptyQueryParameter) Error() string {
	return fmt.Sprintf(
		"The '%s' parameter cannot be empty. Please provide a value for this parameter.", e.name,
	)
}
