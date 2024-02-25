package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 0
		close(c)
	}()

	/*
		comma ok idiom can be used to check if the channel is closed
	*/

	i, ok := <-c
	fmt.Println(i, ok)
	if !ok {
		fmt.Println(i, ", Channel is closed")
	}
	fmt.Println("------------------------")


	i, ok = <-c
	fmt.Println(i, ok)
	if !ok {
		fmt.Println(i, ", Channel is closed")
	}

	fmt.Println("------------------------")
	i, ok = <-c // Not going into deadlock as the channel is closed
	fmt.Println(i, ok)
	if !ok {
		fmt.Println(i, ", Channel is closed")
	}
}
