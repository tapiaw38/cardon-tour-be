package utils

import "time"

func ParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func ToDateString(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format("2006-01-02")
}

func ToInt(r rune) int {
	return int(r - '0')
}

func ToPointer[T any](value T) *T {
	return &value
}
