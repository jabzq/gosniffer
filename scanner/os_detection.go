package scanner

import (
	"fmt"
	"net"
	"time"

	"GoSniffer/config"
)

func DetectOS(host string, port int) *config.OSInfo {
	ttl := estimateTTL(host)
	return detectOSByTTL(ttl)
}

func estimateTTL(host string) int {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:80", host), 2*time.Second)
	if err != nil {
		return 128
	}
	defer conn.Close()
	return 128
}

func detectOSByTTL(ttl int) *config.OSInfo {
	switch {
	case ttl <= 64:
		return &config.OSInfo{Name: "Linux/Unix", Accuracy: 0.6, Family: "Unix-like"}
	case ttl <= 128:
		return &config.OSInfo{Name: "Windows", Accuracy: 0.7, Family: "Windows"}
	default:
		return &config.OSInfo{Name: "Unknown", Accuracy: 0.3, Family: "Unknown"}
	}
}
