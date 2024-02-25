package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("CPUs\t\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())

	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			time.Sleep(100 * time.Microsecond)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("i : ", i, "GoRoutines In Loop \t\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter:\t\t", counter)
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
}
