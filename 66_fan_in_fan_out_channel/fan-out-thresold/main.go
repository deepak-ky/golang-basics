package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println("values of c2 : ", v)
	}

	fmt.Println("about to exit")
}

func populate(c1 chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("putting value into c1 : ", i)
		c1 <- i
	}
	close(c1)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup
	const thresoldGoroutines = 3
	wg.Add(thresoldGoroutines)
	for i := 0; i < thresoldGoroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					fmt.Println("received value from i : ", i, " c1 : ", v2)
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(v2 int) int {
	//time.Sleep(time.Duration(v2) * time.Second)
	return v2
}
