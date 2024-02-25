package main

import "fmt"

func main() {
	c := make(chan int, 1) // bidirectional channel -> can send and receive both
	cr := make(<-chan int) // receive channel -> cannot send values to this channel 	( cr <- 42, not allowed)
	cs := make(chan<- int) // send channel -> cannot receive values from this channel  ( <-cs , not allowed)

	fmt.Printf("c type : %T\n", c)
	fmt.Printf("cr type : %T\n", cr)
	fmt.Printf("cs type : %T\n", cs)

	c <- 42
	fmt.Println(<-c)

	/*

		//cr <- 42 // invalid operation: cannot send to receive-only channel cr (variable of type <-chan int)
		fmt.Println(<-cr)

		cs <- 42
		// fmt.Println(<-cs) // invalid operation: cannot receive from send-only channel cs (variable of type chan<- int)

	*/

	/* 

	// General to specific works, no other case works 

	// Specific to general doesn't assign
	c = cr // cannot use cr (variable of type <-chan int) as chan int value in assignment
	c = cs // cannot use cs (variable of type chan<- int) as chan int value in assignment

	// Specific to Specifc doesn't assign
	cr = cs // cannot use cs (variable of type chan<- int) as <-chan int value in assignment

	// General to specific works
	cr = c
	cs = c 
	*/
	
	
	/* 
	fmt.Println("\n\nType Conversions : ")

	// specific to general type conversion
	// fmt.Printf("cr type now : %T\n", (chan int) (cr)) // cannot convert cr (variable of type <-chan int) to type chan int
	// fmt.Printf("cs type now : %T\n", (chan int) (cs)) // cannot convert cs (variable of type chan<- int) to type chan int 


	// specific to specific does not convert
	// fmt.Printf("cr type now : %T\n", (chan<- int)(cr)) // cannot convert cr (variable of type <-chan int) to type chan<- int

	
	// general to specific type conversion
	fmt.Printf("c type now : %T\n", (<-chan int) (c))
	fmt.Printf("c type now : %T\n", (chan<- int) (c))
	
	*/
}
