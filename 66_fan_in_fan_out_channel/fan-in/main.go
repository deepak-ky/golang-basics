package main

import (
	"fmt"
	"sync"
)

func main() {
	smaller := make(chan int)
	greater := make(chan int)
	fanIn := make(chan int)

	go send(smaller, greater)

	go receive(smaller, greater, fanIn)

	for v := range fanIn {
		fmt.Println("fanIn : ", v)
	}

	fmt.Println("about to exit")
}

func receive(smaller, greater <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range smaller {
			fmt.Println("smaller : ", v)
			fanIn <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range greater {
			fmt.Println("greater : ", v)
			fanIn <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanIn)

	/* for {
		select {
		case v, ok := <-greater:
			if !ok {
				close(fanIn)
				return
			}
			fmt.Println("greater : ", v)
			fanIn <- v
		case v, ok := <-smaller:
			if !ok {
				close(fanIn)
				return
			}
			fmt.Println("smaller : ", v)
			fanIn <- v
		}
	} */

}

func send(smaller, greater chan<- int) {
	for i := 0; i < 10; i++ {
		if i < 5 {
			smaller <- i
		} else {
			greater <- i
		}
	}

	close(smaller)
	close(greater)
}
