package scanner

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"GoSniffer/config"
)

func GrabBanner(host string, port int, timeout time.Duration) string {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		return ""
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(timeout))
	reader := bufio.NewReader(conn)

	switch port {
	case 80, 443, 8080, 8443:
		fmt.Fprintf(conn, "HEAD / HTTP/1.1\r\nHost: %s\r\nUser-Agent: GoSniffer/1.0\r\nConnection: close\r\n\r\n", host)
	case 21:
		// FTP - read banner initial
	case 22:
		// SSH - read banner initial
	case 25, 587:
		fmt.Fprintf(conn, "EHLO example.com\r\n")
	default:
		return readInitialBanner(reader)
	}

	return readBannerResponse(reader, port)
}

func readInitialBanner(reader *bufio.Reader) string {
	banner, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(banner)
}

func readBannerResponse(reader *bufio.Reader, port int) string {
	var response strings.Builder
	maxLines := 10

	for i := 0; i < maxLines; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		response.WriteString(strings.TrimSpace(line) + " | ")

		if (port == 80 || port == 443 || port == 8080) && strings.TrimSpace(line) == "" {
			break
		}
	}

	result := response.String()
	if len(result) > 200 {
		result = result[:200] + "..."
	}
	return result
}

func DetectService(port int, banner string) string {
	if banner != "" {
		bannerUpper := strings.ToUpper(banner)
		switch {
		case strings.Contains(bannerUpper, "HTTP"):
			return "HTTP Server"
		case strings.Contains(bannerUpper, "SSH"):
			return "SSH"
		case strings.Contains(bannerUpper, "FTP"):
			return "FTP Server"
		case strings.Contains(bannerUpper, "SMTP"):
			return "SMTP"
		case strings.Contains(bannerUpper, "MYSQL"):
			return "MySQL"
		case strings.Contains(bannerUpper, "POSTGRES"):
			return "PostgreSQL"
		}
	}

	if service, exists := config.ServiceMap[port]; exists {
		return service
	}

	return "Unknown"
}
