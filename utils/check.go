package utils

import (
	"net"
)

func IPChecker(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	} else {
		return true
	}
}
