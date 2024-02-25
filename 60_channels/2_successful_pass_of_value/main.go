package main

func main() {
	c := make(chan int)
	/*
		This program does run.

		And it runs because, in this program there is a goroutine to receive the value, which is the main original goroutine

		line 20 -> new go routine
		line 23 -> main go routine

		line 20 go routine sends the value to the channel and line 23 goroutine receives the value from the channel "at the same time".

		The passing of baton has happened in this case, at line 20 and line 23

		https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922344#questions/21390794
	*/

		c <- 42


	//fmt.Println(<-c)
}

/*
	This program also works

	https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922344#questions/21390794

	package main
	import "fmt"
	func main() {
		c := make(chan int)
		go func(){
			c <- 42
		}()
	}
*/
