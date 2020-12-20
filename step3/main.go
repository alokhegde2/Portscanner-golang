//In step2 we know how to scan 1024 port using for loop,in step2 we just printing the port number but here we are assigning it to net.Dial
package main

import (
	"fmt"
	"net"
)

func main() {
	var address string
	ports := 1024
	fmt.Println("[*]Enter the target to scan:")
	fmt.Scanf("%s", &address)
	fmt.Println("[*]Enter the number of ports to scan(By default 1024):")
	fmt.Scanf("%d", &ports)
	for i := 1; i <= ports; i++ {
		address := fmt.Sprintf("%s:%d", address, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//port is closed or filtered
			fmt.Println("[+] Scanning Ports...")
			continue
		}
		conn.Close()
		fmt.Printf("%d open \n", i)
	}
}
