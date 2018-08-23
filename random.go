package helper

import (
	"crypto/rand"
	"io"
	"math/big"
	mrand "math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"math"
)

//uuid+unix time
func RandId() string {
	return strconv.FormatInt(int64(uuid.New().Time()/10000000000)*10000000000+Int64Range(100000000, 10000000000), 10)
}

func RandIdInt64() int64 {
	return int64(uuid.New().Time()/10000000000)*10000000000 + Int64Range(100000000, 10000000000)
}

//0123456789 select 6 password number
var RandNumTmps = []byte("0123456789")

func RandNumber(length int, chars []byte) string {
	newPwd := make([]byte, length)
	random := make([]byte, length+(length/4)) // storage for random bytes.
	charsLength := byte(len(chars))
	max := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, random); err != nil {
			panic(err)
		}
		for _, c := range random {
			if c >= max {
				continue
			}
			newPwd[i] = chars[c%charsLength]
			i++
			if i == length {
				return string(newPwd)
			}
		}
	}
	panic("unreachable")
}

//[min,max]
func Int64Range(min, max int64) int64 {
	b, err := rand.Int(rand.Reader, big.NewInt(max+1-min))
	if err != nil {
		return max
	}
	return min + b.Int64()
}

var r = mrand.New(mrand.NewSource(time.Now().UnixNano()))

func RandInt(min, max int) int {
	return r.Intn(max+1-min) + min
}

func RandInt32(min, max int32) int32 {
	return r.Int31n(max+1-min) + min
}

func RandUInt32(min, max uint32) uint32 {
	return uint32(r.Int31n(int32(max+1-min)) + int32(min))
}

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandomLetters(n int, alphabets ...byte) string {
	var bytes = make([]byte, n)
	var randby bool
	if num, err := r.Read(bytes); num != n || err != nil {
		mrand.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[mrand.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[mrand.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return Byte2String(bytes)
}

func GenOrderNo(userId string) string {
	return time.Now().Format("0201150405") + userId
}

func NormFloat64(num int, sq, ar float64) (f []float64) {
	for i := 0; i < num; i++ {
		f = append(f, math.Abs(r.NormFloat64()*sq+ar))
	}
	return
}
