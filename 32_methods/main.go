package main

import "fmt"

/* Methods in go -> attach function to a type
Any value of that type will have access to that function as a method
Implemented with the help of receiver*/

type book struct {
	name string
}

func (b book) printBookName() {
	fmt.Println("The book's name is",b.name)
}

func printBookName2(b book){
	fmt.Println("The book's name is:",b.name)
}

func main() {
	b1 := book{
		name: "A thousand splendid suns",
	}

	b2 := book{
		name: "The almanack of naval ravikant",
	}

	fmt.Println(b1)
	b1.printBookName()
	printBookName2(b1)

	b2.printBookName()
}
