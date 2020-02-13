package main

import (
	"fmt"
	"os"
)

func main() {
	// simpleDemo()
	// closeDeferred()
	scopesDemo()
}

func simpleDemo() {
	fmt.Println("SIMPLE DEMO")
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
	fmt.Println("\nCLOSE DEFERRED")
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

func scopesDemo() {
	fmt.Println("\nSCOPES DEMO")
	defer fmt.Println("Defer 1")
	fmt.Println("Line 1")
	deferFunc2()
	deferFunc3()
	defer fmt.Println("Defer 4")
	fmt.Println("Line 4")
}

func deferFunc2() {
	defer fmt.Println("Defer 2")
	fmt.Println("Line 2")
}

func deferFunc3() {
	defer fmt.Println("Defer 3")
	fmt.Println("Line 3")
}
