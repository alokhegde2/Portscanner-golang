package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:90")
	if err != nil {
		fmt.Println("Connection Succefull")
	}
}
