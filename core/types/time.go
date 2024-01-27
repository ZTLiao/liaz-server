package types

import (
	"core/utils"
	"strconv"
	"strings"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = Time(time.Time{})
		return
	}
	if strings.Contains(string(data), utils.DASHED) && strings.Contains(string(data), utils.DOT) {
		now, err := time.ParseInLocation(`"`+utils.NORM_DATETIME_MS_PATTERN+`"`, string(data), time.Local)
		*t = Time(now)
		return err
	} else if strings.Contains(string(data), utils.DASHED) {
		now, err := time.ParseInLocation(`"`+utils.NORM_DATETIME_PATTERN+`"`, string(data), time.Local)
		*t = Time(now)
		return err
	} else {
		timestamp, err := strconv.ParseInt(string(data), 10, 64)
		*t = Time(time.Unix(timestamp/1000, (timestamp%1000)*int64(time.Millisecond)))
		return err
	}
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(time.Time(t).UnixMilli()), 10)), nil
}
