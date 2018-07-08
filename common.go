package helper

import (
	"math"

	"github.com/jinbanglin/log"
	"github.com/json-iterator/go"
	"time"
)

func Marshal2String(data interface{}) string {
	s, err := jsoniter.MarshalToString(data)
	if err != nil {
		log.Error(" |err ", err)
	}
	return s
}

func Marshal2Bytes(data interface{}) []byte {
	s, err := jsoniter.Marshal(data)
	if err != nil {
		log.Error(" |err ", err)
	}
	return s
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func NextZeroDateWithSub() (time.Time, time.Time) {
	now := time.Now()
	// 计算下一个零点
	next := now.Add(time.Hour * 24)
	next = time.Date(
		next.Year(),
		next.Month(),
		next.Day(),
		0,
		0,
		0,
		0,
		next.Location())
	return now, next
}
