package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var maxBarValue int
var maxFooValue int

func main() {
	maxBarValue = 0
	maxFooValue = 0

	fmt.Println("=> Beginning \t\t")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())

	/* 
		wg.Add(0)

		Value added inside the waitgroup counter is less than number of goroutines that are going to be
		spawned. In this case, when wg.Done() occurs it will decrement the counter from 0 to -1 which will give the following error : 

		panic: sync: negative WaitGroup counter
		goroutine 6 [running]:

		To achieve the above case keep the range of foo till 10
		// -> for i := 0; i < 10; i++ {
	*/

	wg.Add(1)

	/* 
		wg.Add(2)

		The above case is also a issue, we have wg.Done() once in our code, so that will decrement the counter and keep it to 1. 

		The following error would be printed: 
			fatal error: all goroutines are asleep - deadlock!
		
		This is because of a situation where wg.Wait() waits for 2 goroutines to finish, 
			but only one actually finishes, 
			and there is no other goroutine to finsh ("all goroutines are asleep"),
			leaving the program in a deadlock.

			wg.Wait() -> Waiting for something(a goroutine) to finish which does not exist => DEADLOCK
	*/


	go foo()
	bar()

	fmt.Println("=> Intermediate \t\t")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("=> Ending \t\t")
	fmt.Println("maxFooValue\t\t", maxFooValue)
	fmt.Println("maxBarValue\t\t", maxBarValue)
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	/* 
		If waitgroups are not used then the NumGoroutines in the end would still be 2
		and the program would end with 2 go routines running which is not good
	*/
	fmt.Println("Go Routines\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10000; i++ { // foo does not get a chance to completely complete
		fmt.Println("foo: ", i)
		if i > maxFooValue {
			maxFooValue = i
		}
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("bar: ", i)
		if i > maxBarValue {
			maxBarValue = i
		}
	}
}
