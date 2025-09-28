package scanner

import (
	"strconv"
	"strings"
)

func ParsePorts(portsFlag string) []int {
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
