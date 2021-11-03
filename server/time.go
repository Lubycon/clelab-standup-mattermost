package main

import (
	"time"
)

func getNowTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now()
	now.In(loc)

	return time.Now()
}
