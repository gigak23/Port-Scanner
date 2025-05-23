package port

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

// Takes protcol such as TCP, UDP, etc...
// Hostname Ex: localhost
// Port Ex: 80
func ScanPort(protocol, hostname string, port string) *ScanResult {
	address := hostname + ":" + port
	pt := port + "/" + protocol
	// Dial function connects to protocol port
	// Ex: TCP, UDP, HTTP, etc...
	// This is the same fuction but it times out if
	// it takes to long to connect to port
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		return &ScanResult{
			Port:  pt,
			State: "closed",
		}
	}
	defer conn.Close()

	return &ScanResult{
		Port:  pt,
		State: "open",
	}

}

// Scans through a range of ports and checks if they are open
func InitialScan(hostname, protocol string) {
	var results []*ScanResult

	for i := 1; i <= 1024; i++ {
		results = append(results, ScanPort(protocol, hostname, strconv.Itoa(i)))
		fmt.Println(*results[i-1])
	}

}
