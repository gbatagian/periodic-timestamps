package ptlist

import "fmt"

type unsupportedPeriodError struct {
	value string
}

func (e unsupportedPeriodError) Error() string {
	return fmt.Sprintf("unsupported period '%s'", e.value)
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
		"timestamp '%s' does not match the required format %s", e.timestamp, e.format,
	)
}

type emptyQueryParameter struct {
	name string
}

func (e emptyQueryParameter) Error() string {
	return fmt.Sprintf(
		"parameter '%s' cannot be empty, please provide a value", e.name,
	)
}
