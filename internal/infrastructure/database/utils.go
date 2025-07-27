package database

import (
	"time"
)

// timeFromUnix converts Unix timestamp to time.Time
func timeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// timeToUnix converts time.Time to Unix timestamp
func timeToUnix() int64 {
	return time.Now().Unix()
} 