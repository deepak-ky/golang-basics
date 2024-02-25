package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		/*
			If the channel is not closed here using close(c), then the below range loop will result in a deadlock

			This is because the range loop pulls off values of a channel, until the channel is closed.

			When there are no more values on the channel (the channel is closed) it exits the range loop
		*/
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
