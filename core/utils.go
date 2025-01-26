package core

import (
	"net"
	"strings"
)

func resolveHostname(ip string) string {
	hosts, err := net.LookupAddr(ip)
	if err != nil || len(hosts) == 0 {
		return ""
	}
	return strings.TrimSuffix(hosts[0], ".") // Remove trailing dot from hostname
}
