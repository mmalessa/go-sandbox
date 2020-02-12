package main

import "fmt"

// only for defer!

func main() {
	defer someFuncion("Text for some function")
	// someFuncion("Text for some function")
	panic("boom")
}

func someFuncion(msg string) {
	if rec := recover(); rec != nil {
		fmt.Println("Recover:", rec)
	}
	fmt.Println(msg)
}
