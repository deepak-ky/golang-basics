package main

import "fmt"

func main() {
	/*
		THIS PROGRAM DOES NOT RUN

		fatal error: all goroutines are asleep - deadlock!

		CHANNELS BLOCK

		1. "c <- 42"  ⬅ This is an issue, as you are sending something on the channel,
		but there is no one to pull it off (no one to receive it)

		2. Sending data and Receving data on a channel should happen at the same time, as in passing a baton in a relay race

		3. And because there was no one to receive in this case, channel blocked the operation, and deadlock occured. No goroutine to receive.

		4. The vice versa of the above statements is also true, if our programs contain the following two lines
				c := make(chan int)
				fmt.Println(<-c)

			then also we have the same error : fatal error: all goroutines are asleep - deadlock!
			because in this case, it is ready to receive, but no one is ready to send
	*/
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}
