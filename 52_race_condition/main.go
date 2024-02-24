package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUs\t\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			time.Sleep(100 * time.Microsecond) // some go routines set the value set by other go routines
			/* 
				If you increase the time 👆, the counter value would be reaching low final values.
				Time to process a single go routine has increased, reading and setting the values happens less frequently 
			*/

			//time.Sleep(5 * time.Second) // In this case all go routines read 0 and set 1
			//runtime.Gosched() // every go routines reads the value set by previous go routine
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("i : ", i, "GoRoutines In Loop \t\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter:\t\t", counter)
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
}
