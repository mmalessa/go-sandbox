package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	demoNoSync()
	demoWithChannels()
	demoWithOneChannel()
	demoWithWaitGroup()
}

func counter(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(1 * time.Second)
	}
}

func demoNoSync() {
	fmt.Println("DEMO NO SYNC")
	go counter("Go1")
	go counter("Go2")
	go func(msg string) {
		fmt.Println("Hi,", msg)
	}("John")
	counter("Norm")
	time.Sleep(2 * time.Second)
}

func counterWithChannels(msg string, num int, sync chan bool) {
	for i := 1; i <= num; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(1 * time.Second)
	}
	sync <- true
}

func demoWithChannels() {
	fmt.Println("\nDEMO WITH CHANNELS")
	sync1 := make(chan bool)
	sync2 := make(chan bool)
	go counterWithChannels("Go1", 2, sync1)
	go counterWithChannels("Go2", 3, sync2)
	<-sync1
	<-sync2
}

func counterWithOneChannel(msg string, num int, sync chan int) {
	for i := 1; i <= num; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(1 * time.Second)
	}
	sync <- 1
}

func demoWithOneChannel() {
	fmt.Println("\nDEMO WITH ONE CHANNEL")
	sync := make(chan int)
	howMany := 3

	for i := 1; i <= howMany; i++ {
		go counterWithOneChannel(strconv.Itoa(i), i, sync)
	}

	num := 0
	for {
		num = num + <-sync
		if howMany == num {
			fmt.Println("FINISH")
			close(sync)
			break
		}
		fmt.Println("Still working...", num)
	}
}

func counterWithWaitGroup(msg string, wg *sync.WaitGroup) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func demoWithWaitGroup() {
	fmt.Println("\nDEMO WITH WAIT GROUP")
	wg := &sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go counterWithWaitGroup(strconv.Itoa(i), wg)
	}
	wg.Wait()
}
