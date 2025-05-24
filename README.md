# GoSniffer

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/jbzq/Gosniffer/actions/workflows/go.yml/badge.svg)](https://github.com/jbzq/Gosniffer/actions)
[![Codecov](https://codecov.io/gh/jbzq/GoSniffer/branch/main/graph/badge.svg)](https://codecov.io/gh/jbzq/Gosniffer)
[![Go Report Card](https://goreportcard.com/badge/github.com/jbzq/Gosniffer)](https://goreportcard.com/report/github.com/jbzq/Gosniffer)
[![Release](https://img.shields.io/github/v/release/jbxq/Gosniffer)](https://github.com/jbzq/Gosniffer/releases)

> [!NOTE]
>
> Still in the testing phase, 
> only ports 80, 443, 3306, 21, 22, 8080, and 53 are recognized, 
> contributions are welcome
>

## Features 
- **Concurrent scanning** (uses goroutines for speed)
- **Service detection** (HTTP, SSH, FTP, etc.)
- **Lightweight** (single binary, no dependencies)

## Instalation

### Using go

```
$ go install github.com/jbzq/Gosniffer@latest
```

### Using git

```
$ git clone https://github.com/jbzq/Gosniffer.git
$ cd git
$ go build -o gosniffer main.go  # in Unix/Mac
$ go build -o gosniffer.exe  # in Windows
```

## Usage 

```
$ go run main.go -h google.com -p 80,21 -t 2s
```

```
$ ./gosniffer -h google.com -p 80,21 -t 2s
```

### exemple

```
jbz@bird ~/g/GoSniffer (main) [1]> go run main.go -h kali.org -p 1-1000

🔍 Scanning kali.org (Ports: 1-1000)...

[+] Port 80 open (HTTP)
[+] Port 53 open (DNS)
[+] Port 443 open (HTTPS)
```

## License

[Mit License](/LICENSE)
