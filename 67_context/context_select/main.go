package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1 : ", ctx.Err())
	fmt.Println("num goroutines 1 : ", runtime.NumGoroutine())

	go func(){
		n := 0
		for {
			select{
			case <- ctx.Done():
				fmt.Println("ctx.Done() happened", ctx.Done())
				return
			default:
				n++;
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("error check 2 : ", ctx.Err())
	fmt.Println("num goroutines 2 : ", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	/* 
		whenever cancel() is called
		this condition is met..
			case <- ctx.Done():
				fmt.Println("ctx.Done() happened", ctx.Done())
				return
	
	*/
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 8)
	fmt.Println("error check 3 : ", ctx.Err())
	fmt.Println("num goroutines 3 : ", runtime.NumGoroutine())

}