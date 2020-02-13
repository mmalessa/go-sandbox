package main

import (
	"fmt"
	"os"
)

func main() {
	simpleDemo()
	closeDeferred()
	defer someFunc1()
	someFunc2()
}

func simpleDemo() {
	someFunction("Normal 1")
	defer someFunction("Deferred 1")
	defer someFunction("Deferred 2")
	defer someFunction("Deferred 3")
	defer someFunction("Deferred 4")
	someFunction("Normal 2")
}

func someFunction(msg string) {
	fmt.Println(msg)
}

func closeDeferred() {
	f, err := os.Open("/etc/issue")
	if err != nil {
		panic(err)
	}
	// defer f.Close()
	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()
	fmt.Println(*f)
}

func someFunc1() {
	defer fmt.Println("Defer. SomeFunc1")
	fmt.Println("SomeFunc1")
}

func someFunc2() {
	defer fmt.Println("Defer. SomeFunc2")
	fmt.Println("SomeFunc2")
}
