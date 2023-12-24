package main

import "fmt"

type person struct {
	first               string
	last                string
	favIceCreamFlavours []string
}

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

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p1.favIceCreamFlavours)

	for _, favIceCreamflavour := range p1.favIceCreamFlavours {
		fmt.Printf("%s ", favIceCreamflavour)
	}
	fmt.Println()

	fmt.Println(p2)
	fmt.Println(p2.first, p2.last)
	fmt.Println(p2.favIceCreamFlavours)

	for _, favIceCreamflavour := range p2.favIceCreamFlavours {
		fmt.Printf("%s ", favIceCreamflavour)
	}
}
