package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go fooNew()
	bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func fooNew() {
	for i := 0; i < 1000; i++ {
		fmt.Println("fooNew: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("bar: ", i)
	}
}
