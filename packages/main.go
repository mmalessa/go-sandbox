package main

import (
	"fmt"

	"packages/server"
)

func main() {
	fmt.Println("START")

	s := server.MyServer{
		Address: "localhost",
		Name:    "myserver",
	}
	fmt.Println(s)
}
