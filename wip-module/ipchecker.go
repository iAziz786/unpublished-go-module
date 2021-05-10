package ipchecker

import (
	"net"
)

func IsValidIP(host string) bool {
	if ip := net.ParseIP(host); ip != nil {
		return true
	}

	return false
}
