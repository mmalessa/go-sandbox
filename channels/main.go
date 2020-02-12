package main

/*
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
*/
import (
	"fmt"
	"time"
)

func main() {
	demoUnbuffered()
	demoBuffered()
	demoOneDirection()
	demoWithSelect()
}

func demoUnbuffered() {
	fmt.Println("DEMO UNBUFFERED")
	chInt := make(chan int)
	chStr := make(chan string)

	go func() {
		chInt <- 123
		chStr <- "some text"
	}()

	fmt.Println(<-chInt)
	fmt.Println(<-chStr)
	close(chInt)
	close(chStr)
}

func demoBuffered() {
	fmt.Println("\nDEMO BUFFERED")
	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	// ch <- "d"
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}

func demoOneDirection() {
	fmt.Println("\nDEMO ONE DIRECTION")
	chToReceive := make(chan string, 1)
	chToSend := make(chan string, 1)

	go receiveOnly(chToReceive)
	go sendOnly(chToSend)
	chToReceive <- "some text 1"
	fmt.Println(<-chToSend)

	time.Sleep(1 * time.Second)
	close(chToReceive)
	close(chToSend)

}

func receiveOnly(c <-chan string) {
	fmt.Println(<-c)
}

func sendOnly(c chan<- string) {
	c <- "some text 2"
}

func demoWithSelect() {
	fmt.Println("\nDEMO WITH SELECT")
	chInt := make(chan int, 2)
	chStr := make(chan string, 2)

	chInt <- 10
	chInt <- 20
	chStr <- "first hello"
	chStr <- "second hello"

	msgsNum := len(chInt) + len(chStr)
	for i := 0; i < msgsNum; i++ {
		select {
		case num := <-chInt:
			fmt.Println("Num:", num)
		case str := <-chStr:
			fmt.Println("String:", str)
		default:
			fmt.Println("default nothing")
		}
	}
	close(chInt)
	close(chStr)
}
