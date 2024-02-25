package main

import (
	"fmt"
	"time"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send on the above channels
	go send(even, odd, quit)

	// receive from the above channels
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Received from even channel : ",v)
		case v := <-odd:
			fmt.Println("Received from odd  channel : ", v)
		case v := <-quit :
			fmt.Println("Received from quit channel : ", v)
			return
		// default: 
		// 	fmt.Println("Waiting ⌛")
		}
	}
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 5; i++ {
		if i % 2 == 0 {
			time.Sleep(1 * time.Second)
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- -1
}