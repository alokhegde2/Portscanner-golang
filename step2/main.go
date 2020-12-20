//here we are finding open port from 1 to 1024
package main

import "fmt"

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.namp.org:%d", i) // value of i in integer we have to convert it in to string for that we are using fmt.Sprintf and the result is assigned to address
		fmt.Println(address)                            //here we are printing the address which is assigned above
	}
}
