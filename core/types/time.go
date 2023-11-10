package types

import (
	"core/utils"
	"strconv"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+utils.DEFAULT_DATETIME_PATTERN+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).UnixMilli(), 10)), nil
}
