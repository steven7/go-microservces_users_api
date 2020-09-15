package date_utils

import "time"

const (
	dateLayout     =  "2006-01-02T15:04:05Z"
	date_Db_Layout = "2006-01-02 15:04:05"
)

func GetCurrentDateTime() time.Time {
	return time.Now().UTC()
}

func GetCurrentDateTimeString() string {
	return GetCurrentDateTime().Format(dateLayout)
}

func GetCurrentDateTimeDBFormat() string {
	return GetCurrentDateTime().Format(date_Db_Layout)
}