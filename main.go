package main

import (
	"flag"
	"fmt"
	"time"

	"GoSniffer/config"
	"GoSniffer/scanner"
)

func printBanner() {
	banner := `
   ___                _  __  __
  / _ \___  ___ _ __ (_)/ _|/ _| ___ _ __
 / /_\/ _ \/ __| '_ \| | |_| |_ / _ \ '__|
/ /_\\ (_) \__ \ | | | |  _|  _|  __/ |
\____/\___/|___/_| |_|_|_| |_|  \___|_|
`
	fmt.Printf("%s\n", banner)
}

func main() {
	printBanner()

	host := flag.String("h", "example.com", "Target Host")
	portsFlag := flag.String("p", "80,443", "Ports (ex: '1-1000' or '22,80,443')")
	timeout := flag.Duration("t", 2*time.Second, "Timeout for each connection")
	scanType := flag.String("s", "connect", "Scan type (connect, syn)")
	banner := flag.Bool("b", false, "Enable banner grabbing")
	vulnCheck := flag.Bool("v", false, "Enable vulnerability checks")
	osDetection := flag.Bool("os", false, "Enable OS detection")
	output := flag.String("o", "text", "Output format")

	flag.Parse()

	scanConfig := config.ScanConfig{
		Host:        *host,
		Ports:       *portsFlag,
		Timeout:     *timeout,
		ScanType:    *scanType,
		Banner:      *banner,
		VulnCheck:   *vulnCheck,
		OSDetection: *osDetection,
		Output:      *output,
	}

	scanner.RunScan(scanConfig)
}
