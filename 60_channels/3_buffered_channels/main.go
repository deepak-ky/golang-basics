package main

import "fmt"

func main() {
	/*
		Simple channels or unbuffered channels do not allow the sending of a value to a channel if there isn't a goroutine available to receive at the same time.

		Buffered channels allow this behavour, you can send values, upto a certain buffer limit, even if there is no goroutine availabe to receive it

		Buffered channels generally employ FIFO within the buffer, see : // Works-3
	*/

	/* // Works-1
	ch := make(chan int, 1)
	ch <- 32
	fmt.Println(<-ch) */

	/* // Deadlock - 1
	ch := make(chan int, 1)
	ch <- 32
	fmt.Println(<-ch)
	fmt.Println(<-ch) */

	/* // Deadlock - 2
	ch := make(chan int, 1)
	ch <- 32
	ch <- 33
	fmt.Println(<-ch) */

	/* // Works-2
	ch := make(chan int, 1)
	ch <- 32
	fmt.Println(<-ch)
	ch <- 33
	fmt.Println(<-ch) */

	// Works-3
	ch := make(chan int, 2)
	ch <- 32
	ch <- 33
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
