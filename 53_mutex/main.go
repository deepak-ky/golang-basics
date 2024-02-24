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
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			/* 
				When the goroutine at i = 1 attempts to lock the mutex (mu.Lock())
   					a. and finds it already locked by the goroutine at i = 0,
   					b. it will indeed wait until the mutex is unlocked by the goroutine at i = 0.
   					c. This is the purpose of using a mutex - to ensure that only one goroutine can access the critical section (in this case, accessing/modifying counter) at any given time.
			*/
			mu.Lock()
			v := counter
			time.Sleep(100 * time.Microsecond)
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("i : ", i, "GoRoutines In Loop \t\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter:\t\t", counter)
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
}
