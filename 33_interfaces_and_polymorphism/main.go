package main

import "fmt"

/*
In Go, Interfaces allow us to have polymorphism
An Interface in go defines a set of method signatures
Polymorphism is the ability of a VALUE of a certain TYPE to also be of another TYPE
In Go, values can be of more than one type
*/

type person struct {
	name string
}

type secretAgent struct {
	person
	liceneseToKill bool
}

func (p person) Introduce() {
	fmt.Println("Hi, I am a person and my name is:", p.name)
}

func (sa secretAgent) Introduce() {
	fmt.Println("Hi, I am a secretAgent and my name is:", sa.name)
}

/*
The below definition means, any value that is of a certain type
which as a method Introduce()
will also be of type human
	1. Any value of type person will also be of type human
	2. Any value of type secretAgent will also be of type human

So, a value can be of more than one type -> Polymorphism
*/
type human interface {
	Introduce()
}

func IntroduceYourself(h human) {
	h.Introduce()
}

func main() {
	p1 := person{
		name: "Deepak Kumar",
	}

	sa1 := secretAgent{
		person: person{
			name: "KD_4",
		},
		liceneseToKill: true,
	}

	fmt.Println(p1)
	fmt.Println(sa1)
	p1.Introduce()
	/* if secretAgent struct does not have a Introduce method, person struct's
	Introduce method gets inner field promotion.*/
	sa1.Introduce()

	fmt.Println("---------------------------------------------")

	IntroduceYourself(p1)
	IntroduceYourself(sa1)


}
