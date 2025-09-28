package scanner

import (
	"strings"
)

var CommonVulns = map[int][]string{
	21:   {"FTP Anonymous login possible", "unencrypted transfer"},
	23:   {"Telnet - unencrypted communication"},
	135:  {"RPC - possible enumeration vulnaribilities"},
	139:  {"NetBIOS - system information possible"},
	445:  {"SMB - verify EternalBlue"},
	1433: {"SQL Server - verify weak credentials"},
	1434: {"SQL Server - UDP enumeration"},
	3389: {"RDP - verify BlueKeep"},
}

func CheckVulnerabilities(port int, service string, banner string) []string {
	var vulns []string

	if portVulns, exists := CommonVulns[port]; exists {
		vulns = append(vulns, portVulns...)
	}

	bannerUpper := strings.ToUpper(banner)

	switch service {
	case "FTP", "FTP Server":
		vulns = append(vulns, "check anonymous connections")
		if strings.Contains(bannerUpper, "VSFTPD") && strings.Contains(bannerUpper, "2.3.4") {
			vulns = append(vulns, "VSFTPD 2.3.4 - Backdoor vulnerability known")
		}

	case "SSH":
		vulns = append(vulns, "check SSH version and weak keys")

	case "HTTP", "HTTPS", "HTTP Server":
		vulns = append(vulns, "Scan web recommended (directors, technologies)")

		if strings.Contains(bannerUpper, "APACHE") {
			vulns = append(vulns, "Apache - check version of vulnaribilities")
		}
		if strings.Contains(bannerUpper, "NGINX") {
			vulns = append(vulns, "Nginx - check settings")
		}

	case "SMTP":
		vulns = append(vulns, "check open relay and valid users")

	case "MySQL", "PostgreSQL":
		vulns = append(vulns, "check standard credentials and permissions")
	}

	return vulns
}
