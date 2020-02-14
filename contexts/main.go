package main

import (
	"context"
	"fmt"
	"time"
)

// based on http://jgao.io/
func main() {
	contextWithTimeout()
	fmt.Println("")
	contextWithDeadline()
	fmt.Println("")
	contextWithCancel()
	fmt.Println("")
	contextWithValue()
}

func contextWithTimeout() {
	fmt.Println("CONTEXT WITH TIMEOUT")
	start := time.Now()
	ch := make(chan struct{})

	// top level context
	ctx := context.Background()

	// context to pass to function
	contextWithTimeout, cancelTimeout := context.WithTimeout(ctx, 1000*time.Millisecond)
	go computation(contextWithTimeout, 2, "Timeout", ch)
	defer cancelTimeout()

	select {
	case val := <-ch:
		fmt.Println("Got result: ", val)
	case <-contextWithTimeout.Done():
		fmt.Println("Computation cancelled by timeout.")
	}

	elapsed := time.Since(start)
	fmt.Printf("Program finished... It took %s \n", elapsed)
}

func contextWithDeadline() {
	fmt.Println("CONTEXT WITH DEADLINE")
	start := time.Now()
	ch := make(chan struct{})

	// top level context
	ctx := context.Background()

	// context to pass to function
	contextWithDeadline, cancelDeadline := context.WithDeadline(ctx, time.Now().Add(1*time.Second))
	go computation(contextWithDeadline, 2, "Timeout", ch)
	defer cancelDeadline()

	select {
	case val := <-ch:
		fmt.Println("Got result: ", val)
	case <-contextWithDeadline.Done():
		fmt.Println("Computation cancelled by deadline.")
	}

	elapsed := time.Since(start)
	fmt.Printf("Program finished... It took %s \n", elapsed)
}

func contextWithCancel() {
	fmt.Println("CONTEXT WITH DEADLINE")
	start := time.Now()
	ch := make(chan struct{})

	// top level context
	ctx := context.Background()

	// context to pass to function
	contextWithCancel, cancel := context.WithCancel(ctx)
	go computation(contextWithCancel, 2, "Timeout", ch)
	defer cancel()
	time.Sleep(time.Duration(1) * time.Second)
	cancel()

	select {
	case val := <-ch:
		fmt.Println("Got result: ", val)
	case <-contextWithCancel.Done():
		fmt.Println("Computation cancelled by cancel.")
	}

	elapsed := time.Since(start)
	fmt.Printf("Program finished... It took %s \n", elapsed)
}

func contextWithValue() {
	fmt.Println("CONTEXT WITH DEADLINE")
	fmt.Println("...TODO")
	// TODO
}

func computation(ctx context.Context, seconds int, name string, c chan struct{}) {
	fmt.Printf("%s: computing started...\n", name)
	time.Sleep(time.Duration(seconds) * time.Second)
	c <- struct{}{}
	fmt.Printf("%s: computing finished...\n", name)
}
