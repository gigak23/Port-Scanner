package main

import (
	"fmt"

	"github.com/gigak23/pscanner/port"
)

func main() {

	fmt.Println("Port scanner in Go")
	port.InitialScan("localhost", "tcp")
}
