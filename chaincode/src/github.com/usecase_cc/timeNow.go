package main

import (
	"time"
	"strconv"
)

func timeNow()(string) {
	sec := time.Now().Unix()      // number of seconds since January 1, 1970 UTC
	s := strconv.FormatInt(sec, 10)
	return s
}