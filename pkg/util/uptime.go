package util

import "time"

var startTime time.Time

func Uptime() time.Duration {
	return time.Since(startTime)
}