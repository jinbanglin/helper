package helper

import (
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetPwd() string {
	s,err:=os.Getwd()
	if err!=nil{
		panic(err)
	}
	return s
}
