package server

import "fmt"

type MyServer struct {
	Address string
	Name    string
}

func init() {
	fmt.Println("server init() called")
}
