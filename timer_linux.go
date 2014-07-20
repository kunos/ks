package ks

import (
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func GetTimeMillis() uint64 {
	n := time.Now()

	return uint64(n.Sub(startTime).Nanoseconds() / 1000000)
}
