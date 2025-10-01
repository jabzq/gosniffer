# GoSniffer

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/jabzq/Gosniffer)](https://goreportcard.com/report/github.com/jabzq/Gosniffer)

> [!NOTE]
> 
> **Now with Advanced Features!** 🚀
> GoSniffer has evolved into a comprehensive network scanning tool with enhanced capabilities while maintaining its lightweight nature.


## Installation

### Using Go
```bash
go install github.com/jabzq/Gosniffer@latest
```

### Using Git 
```bash
git clone https://github.com/jabzq/Gosniffer.git
cd Gosniffer
go build -o gosniffer main.go        # Unix/Linux/Mac
go build -o gosniffer.exe            # Windows
```

### Usage

#### Basic scan
```bash
go run main.go -h google.com -p 80,443 -t 2s
```

#### Advanced Scan with All Features
```bash
./gosniffer -h target.com -p 1-1000 -b -v -os
```

#### Full Port Range with Security Checks
```bash
./gosniffer -h example.com -p 21,22,23,25,53,80,110,143,443,993,995,3389 -b -v -os
```

### Examples

#### Basic Service Detection
```bash
$ go run main.go -h kali.org -p 1-1000

Scanning kali.org (Ports: 1-1000)...

[+] Port 80 open (HTTP)
[+] Port 53 open (DNS) 
[+] Port 443 open (HTTPS)
```

#### Advanced Scan with All Features
```bash
$ ./gosniffer -h uol.com -p 80,443 -b -v -os

Scanning uol.com (Ports: 80,443)...
Mode: connect | Banner: true | VulnCheck: true | OS Detection: true

[+] Port 80 open (HTTP Server)
    └── Banner: HTTP/1.1 301 Moved Permanently | Server: CloudFront | Date: ...
    └── OS: Windows (70% accuracy)
    └── Vulnerabilities: Web scan recommended (directories, technologies)
[+] Port 443 open (HTTP Server)
    └── Banner: HTTP/1.1 400 Bad Request | Server: CloudFront | Date: ...
    └── OS: Windows (70% accuracy)
    └── Vulnerabilities: Web scan recommended (directories, technologies)

Scan completed in 227ms
2 ports scanned, 2 open ports found
```

## Command Line Options

| Flag | Description | Default |
|------|-------------|---------|
| `-h` | Target hostname or IP address | `example.com` |
| `-p` | Ports to scan (comma-separated or range) | `80,443` |
| `-t` | Timeout for each connection | `2s` |
| `-b` | Enable banner grabbing | `false` |
| `-v` | Enable vulnerability checks | `false` |
| `-os` | Enable OS detection | `false` |
| `-s` | Scan type (connect, syn) | `connect` |
| `-o` | Output format (text, json) | `text` |

## Supported Services

The scanner now recognizes services on these common ports:

- **Web Services**: 80 (HTTP), 443 (HTTPS), 8080 (HTTP-Alt), 8443 (HTTPS-Alt)
- **Email**: 25 (SMTP), 110 (POP3), 143 (IMAP), 993 (IMAPS), 995 (POP3S)
- **Remote Access**: 22 (SSH), 3389 (RDP)
- **File Transfer**: 21 (FTP)
- **Databases**: 3306 (MySQL), 5432 (PostgreSQL), 27017 (MongoDB), 6379 (Redis)
- **Network Services**: 53 (DNS)
- **And many more...**

## Project Structure

```bash
GoSniffer/
├── main.go # CLI interface and coordination
├── config/
│ └── config.go # Configuration structures and service mappings
├── scanner/
│ ├── scanner.go # Main scanning logic
│ ├── port_parser.go # Port range parsing
│ ├── banner.go # Banner grabbing functionality
│ ├── vulnerabilities.go # Security checks
│ └── os_detection.go # OS fingerprinting
└── go.mod # dependencies
```

## Contributing

We welcome contributions! The project is actively developed and these features are planned:

- [ ] Raw socket implementation for accurate OS detection
- [ ] JSON output format for integration with other tools
- [ ] Web interface version
- [ ] Distributed scanning capabilities
- [ ] Additional protocol support

Please feel free to submit issues and pull requests!

## License

[MIT License](/LICENSE)

---

**GoSniffer** - Fast, lightweight, and powerful port scanner written in Go. Perfect for network reconnaissance and security assessments.
