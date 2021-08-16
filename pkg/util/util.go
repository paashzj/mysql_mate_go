package util

import (
	"github.com/beego/beego/v2/core/logs"
	"os"
	"strconv"
	"time"
)

func GetEnvStr(key string, value string) string {
	aux := os.Getenv(key)
	if aux != "" {
		return aux
	}
	return value
}

func GetEnvBool(key string, value bool) bool {
	aux := os.Getenv(key)
	if aux != "" {
		return aux == "true"
	}
	return value
}

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
