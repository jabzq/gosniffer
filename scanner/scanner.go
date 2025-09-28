package scanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"GoSniffer/config"
)

func RunScan(scanConfig config.ScanConfig) []string {
	ports := ParsePorts(scanConfig.Ports)
	var results []string
	var wg sync.WaitGroup
	resultChan := make(chan string)

	fmt.Printf("\n🔍 Scanning %s (Ports: %s)...\n", scanConfig.Host, scanConfig.Ports)
	fmt.Printf("📊 Mode: %s | Banner: %v | VulnCheck: %v | OS Detection: %v\n\n",
		scanConfig.ScanType, scanConfig.Banner, scanConfig.VulnCheck, scanConfig.OSDetection)

	startTime := time.Now()

	for _, port := range ports {
		wg.Add(1)
		go scanPort(scanConfig.Host, port, scanConfig, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var openPorts int
	for result := range resultChan {
		results = append(results, result)
		fmt.Println(result)
		openPorts++
	}

	duration := time.Since(startTime)
	stats := fmt.Sprintf("\n📈 Scan completed in %v\n", duration)
	stats += fmt.Sprintf("📊 %d ports scanned, %d open ports found\n", len(ports), openPorts)

	results = append(results, stats)
	fmt.Println(stats)

	return results
}

func scanPort(host string, port int, scanConfig config.ScanConfig, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, scanConfig.Timeout)

	if err != nil {
		return
	}
	defer conn.Close()

	var service string
	var banner string
	var vulnerabilities []string
	var osInfo *config.OSInfo

	if scanConfig.Banner {
		banner = GrabBanner(host, port, scanConfig.Timeout)
		service = DetectService(port, banner)
	} else {
		service = config.ServiceMap[port]
		if service == "" {
			service = "Unknown"
		}
	}

	if scanConfig.OSDetection {
		osInfo = DetectOS(host, port)
	}

	if scanConfig.VulnCheck {
		vulnerabilities = CheckVulnerabilities(port, service, banner)
	}

	result := fmt.Sprintf("\033[32m[+] Port %d open (%s)\033[0m", port, service)

	if banner != "" && scanConfig.Banner {
		result += fmt.Sprintf("\n    └── Banner: %s", banner)
	}

	if osInfo != nil && scanConfig.OSDetection && osInfo.Accuracy > 0.3 {
		accuracyColor := "\033[33m"
		if osInfo.Accuracy > 0.7 {
			accuracyColor = "\033[32m"
		}
		result += fmt.Sprintf("\n    └── \033[36mOS: %s %s(%.0f%% accuracy)\033[0m",
			osInfo.Name, accuracyColor, osInfo.Accuracy*100)
	}

	if len(vulnerabilities) > 0 && scanConfig.VulnCheck {
		result += fmt.Sprintf("\n    └── \033[33mVulnerabilities: %s\033[0m", strings.Join(vulnerabilities, ", "))
	}

	results <- result
}
