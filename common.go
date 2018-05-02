package helper

import (
	"math"

	"github.com/jinbanglin/moss/log"
	"github.com/json-iterator/go"
)

//some constant forever never change
const CONST_TIME_LAYOUT = "2006-01-02"
const CONST_TIME_LAYOUT_COMPLETE = "2006-02-01 15:04:05.000"

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


//保留小数位
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
