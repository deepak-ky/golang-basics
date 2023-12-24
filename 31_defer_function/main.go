package main

import "fmt"

/* A "defer" statement invokes a function whose execution is deferred to the moment the surrounding
	function returns, either because 
	a. the surrounding function exectuted a return statement
	b. reached the end of its function body
	c. or because the corressponding goroutine is panicking

	It's an unusual but effective way to deal with situations such as resources that must be released
	regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file
 */

func main() {
	defer bar()
	foo()	
}

func foo() {
	fmt.Println("I am foo")
}


func bar() {
	fmt.Println("I am bar, I was called before foo, but ran after it 😉 (deferred)")
}