package main

import (
	"fmt"
)

func main() {
	server := newAPIServer(":3000")
	server.Run()
	fmt.Println("Hello, World!")
}
