package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Port and service mapping
var serviceMap = map[int]string{
	80:   "HTTP",
	443:  "HTTPS",
	21:   "FTP",
	22:   "SSH",
	8080: "HTTP-Alt",
	3306: "MySQL",
	53: "DNS",
}

func scanPort(host string, port int, timeout time.Duration, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)

	if err == nil {
		conn.Close()
		service := serviceMap[port]
		results <- fmt.Sprintf("\033[32m[+] Port %d open (%s)\033[0m", port, service)
	}
}

func parsePorts(portsFlag string) []int {
	var ports []int

	if strings.Contains(portsFlag, ",") {
		for _, p := range strings.Split(portsFlag, ",") {
			port, err := strconv.Atoi(p)
			if err == nil {
				ports = append(ports, port)
			}
		}
		return ports
	}

	if strings.Contains(portsFlag, "-") {
		parts := strings.Split(portsFlag, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}
		return ports
	}

	port, err := strconv.Atoi(portsFlag)
	if err == nil {
		ports = append(ports, port)
	}
	return ports
}

func main() {
	host := flag.String("h", "example.com", "Target Host (ex: google.com)")
	portsFlag := flag.String("p", "80,443", "Ports (ex: '1-1000' or '22,80,443')")
	timeout := flag.Duration("t", 2*time.Second, "Timeout for each connection (ex: 1s, 500ms)")
	flag.Parse()

	// Parseando ports
	ports := parsePorts(*portsFlag)

	// Starting WaitGroup and channel results
	var wg sync.WaitGroup
	results := make(chan string)

	fmt.Printf("\n🔍 Scanning %s (Ports: %s)...\n\n", *host, *portsFlag)

	// Starting goroutines for each port
	for _, port := range ports {
		wg.Add(1)
		go scanPort(*host, port, *timeout, &wg, results)
	}

	// Blocking the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
