package config

import "time"

type ScanConfig struct {
	Host        string
	Ports       string
	Timeout     time.Duration
	ScanType    string
	Banner      bool
	VulnCheck   bool
	OSDetection bool
	Output      string
}

type OSInfo struct {
	Name     string
	Accuracy float64
	Version  string
	Family   string
}

type PortResult struct {
	Port            int
	Service         string
	Banner          string
	Vulnerabilities []string
	OSInfo          *OSInfo
}

var ServiceMap = map[int]string{
	80:    "HTTP",
	443:   "HTTPS",
	21:    "FTP",
	22:    "SSH",
	8080:  "HTTP-Alt",
	3306:  "MySQL",
	53:    "DNS",
	25:    "SMTP",
	110:   "POP3",
	143:   "IMAP",
	993:   "IMAPS",
	995:   "POP3S",
	3389:  "RDP",
	5432:  "PostgreSQL",
	27017: "MongoDB",
	6379:  "Redis",
}
