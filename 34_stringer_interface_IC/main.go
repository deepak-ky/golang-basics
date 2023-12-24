package main

import "fmt"

/*
Can adjust defaulting printing for a type

type Stringer interface {
    String() string
}

A Stringer is a type that can describe itself as a string.
The fmt package (and many others) look for this interface to print values
*/

type book struct{
	title string
}

func (b book) String() string {
	return fmt.Sprintf("The book title is %s",b.title)
}

type count int 

func (c count) String() string {
	return fmt.Sprintf("The value of count is %v", int(c)) // c typecasting required here
}

func main(){
	var b1 book
	b1.title = "the bitcoin standard"
	fmt.Println(b1.title) 
	fmt.Println(b1.String())
	// the struct is not printing here, println is internally using String() function 
	// defined on b
	fmt.Println(b1)

	var c count = 42
	fmt.Println(c.String())
	// Println is behaving differently for book type and c type
	// because internally String() function is differently implemented
	fmt.Println(c)
}