package util

import (
	"github.com/beego/beego/v2/core/logs"
	"strconv"
	"time"
)

func ConvStr2Float(str string) float64 {
	if str == "" {
		return 0
	}
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		logs.Warn(err)
		return 0
	}
	return result
}

func Schedule(delay int, duration time.Duration, f func()) {
	go func() {
		for {
			f()
			time.Sleep(time.Duration(delay) * duration)
		}
	}()
}
