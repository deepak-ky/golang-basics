package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Deepak",
		last:  "Kumar",
		age:   24,
	}

	fmt.Println(p1)
	//to print the type and detailed type
	fmt.Printf("%T \t %#v\n", p1, p1)

	//access to field names
	fmt.Println(p1.first, p1.last, p1.age)
}
