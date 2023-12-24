package main

import (
	"fmt"
	"log"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprintf("The book title is %s", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprintf("The value of count is %v", int(c)) // c typecasting required here
}

// Anything can be of type Stringer, it just needs to implement String() function
func logInfo(v fmt.Stringer) {
	log.Println("Module 138", v.String()) // v would also work
}

func main() {
	b1 := book{
		title: "How to win friends and influence people",
	}

	var pages count = 300 

	//fmt.Println internally uses stringer interface
	fmt.Println(b1)    
	fmt.Println(pages)

	//log.Println internally uses stringer interface
	log.Println(b1)
	log.Println(pages)

	//logInfo can be passed with anything which has a type stringer
	logInfo(b1)
	logInfo(pages)

	w := 34
	fmt.Println(w) // does this not use String() internally ??
	//logInfo(w) // w is of type int, int does not implement String(), hence it cannot be passed to 
	// logInfo function
	/* 
	.\main.go:48:10: cannot use w (variable of type int) as fmt.Stringer value in argument to logInfo: 
	int does not implement fmt.Stringer (missing method String)
	*/
}

