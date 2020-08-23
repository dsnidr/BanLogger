package helpers

import "time"

// GetTimestamp takes in a unix timestamp (int64) and returns a formatted datetimestamp
func GetTimestamp(unixtime int64) string {
	return time.Unix(unixtime, 0).Format("02/01/06 03:04:05 PM")
}
