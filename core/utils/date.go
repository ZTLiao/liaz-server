package utils

import "time"

const (
	NORM_MONTH_PATTERN       = "2006-01"
	NORM_DATE_PATTERN        = "2006-01-02"
	NORM_DATETIME_PATTERN    = "2006-01-02 15:04:05"
	NORM_DATETIME_MS_PATTERN = "2006-01-02 15:04:05.999"
)

func GetStartOfWeek(dateTime time.Time) time.Time {
	return dateTime.AddDate(0, 0, -int(dateTime.Weekday())+1)
}

func GetEndOfWeek(dateTime time.Time) time.Time {
	return GetStartOfWeek(dateTime).AddDate(0, 0, 6)
}
