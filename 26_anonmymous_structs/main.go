package main

import "fmt"

type book struct {
	name   string
	author string
	pages  int
}

func main() {
	b1 := book{
		name:   "Atmoic Habits",
		author: "James Clear",
		pages:  300,
	}

	// we do not use the defined struct book here
	// can be used in code when we just want a struct at one place and do not 
	// want to defince a complete struct for it
	b2 := struct {
		name   string
		author string
		pages  int
	}{
		name:   "The Psychology of Money",
		author: "Morgan Housel",
		pages:  250,
	}

	fmt.Println(b1)
	fmt.Println(b2)

	fmt.Printf("b1 type : %T\n",b1) //b1 type : main.book
	fmt.Printf("b2 type : %T",b2)	//b2 type : struct { name string; author string; pages int }
}