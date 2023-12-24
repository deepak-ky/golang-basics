package main

import "fmt"

func main() {

	f1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Deepak",
		friends: map[string]int{
			"arshad": 3,
			"jappi":  2,
			"dj":     1,
		},
		favDrinks: []string{"gin", "water", "martini"},
	}

	fmt.Println(f1)
	fmt.Println(f1.first)
	fmt.Println(f1.friends)
	fmt.Println(f1.friends["dj"])
	fmt.Println(f1.friends["dj-jais"])

	for k,v := range f1.friends {
		fmt.Println(f1.first, "friends with",k,"(avgDrinkCount",v,")")
	}
	fmt.Println(f1.favDrinks)
	fmt.Println(f1.favDrinks[2])
	//fmt.Println(f1.favDrinks[3]) // index out of range
}
