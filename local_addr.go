package helper

import (
	"net"
	"strings"
)

func LocalAddr() string {
	ifaces, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, addr := range ifaces {
		ret := addr.String()
		if strings.Contains(ret, "10.32") {
			continue
		}
		if strings.Contains(ret, "127.") {
			continue
		}
		li := strings.LastIndex(ret, "/")
		return ret[0:li]
	}
	return "unknown"
}
