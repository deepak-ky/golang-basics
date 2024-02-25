package main

import (
	"fmt"
)

//var wg sync.WaitGroup

func main() {
	c := make(chan int)

	//wg.Add(2)

	// send
	go foo(c)

	// receive
	// go bar(c)
	bar(c)
	/* 
		To execute this program without the help of wait groups, one thing that you can do is don't launch function bar in a goroutine, 
		
		what this would help is it will stop the control flow of the program until it pulls the value from the channel.

		just have bar(c)
	*/

	//wg.Wait()

	fmt.Println("about to exit")
}

func foo(c chan<- int) {
	c <- 14
	//wg.Done()
}

func bar(c <-chan int) {
	fmt.Println(<-c)
	//wg.Done()
}
