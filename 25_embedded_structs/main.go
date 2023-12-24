package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licenseToKill   bool
	first string
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "Deepak",
			last:  "Kumar",
			age:   23,
		},
		licenseToKill: true,
		first: "kd_4",
	}

	fmt.Println(sa1)
	fmt.Println(sa1.person)
	fmt.Println(sa1.licenseToKill)

/* 	even though secretAgent does not have a field last, I am still able to 
	access the field due to inner field promotion. This behaviour would not have 
	been possible if we had given the person field in secretAgent some variable name
	for .eg  
	type secretAgent struct {
		personalDetails person
		licenseToKill   bool
		first string
	}
*/
	fmt.Println(sa1.last) 

	//both the secretAgent and person structs have the field first, so when I do
	fmt.Println(sa1.first)
	//it prints the value of first field in secretAgent struct
	//to print first field in person struct, we need to specify
	fmt.Println(sa1.person.first)

}
