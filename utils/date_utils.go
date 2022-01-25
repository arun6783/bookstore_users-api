package dateutils

import (
	"time"
)

const (
	apiDateFormat = "2016-12-22T15:30:30Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}
