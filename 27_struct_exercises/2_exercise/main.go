package main

import "fmt"

type person struct {
	first               string
	last                string
	favIceCreamFlavours []string
}

//var lastNamePersonMapping map[string]person = make(map[string]person)
var lastNamePersonMapping = make(map[string]person)

func main() {
	p1 := person{
		first: "Deepak",
		last:  "Kumar",
		favIceCreamFlavours: []string{
			"vanilla", "butterscotch",
		},
	}

	p2 := person{
		first: "KD",
		last:  "4",
		favIceCreamFlavours: []string{
			"pista", "strawberry",
		},
	}

	//fmt.Println(lastNamePersonMapping)
	lastNamePersonMapping[p1.last] = p1
	lastNamePersonMapping[p2.last] = p2
	fmt.Println(lastNamePersonMapping)	

	for _ , p := range lastNamePersonMapping{
		fmt.Printf("favourite ice-cream flavours for %s => ",p.first)
		for _, favIceCreamflavour := range p.favIceCreamFlavours {
			fmt.Printf("%s ", favIceCreamflavour)
		}
		fmt.Println()
	}
}
