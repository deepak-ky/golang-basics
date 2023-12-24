package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int16
	color string
}

func main() {

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "tesla",
		model: "cybertruck",
		doors: 4,
		color: "grey",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "hyundai",
		model: "aura",
		doors: 4,
		color: v1.color,
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.make, v1.model, v1.electric, v1.color)
	fmt.Println(v2.make, v2.model, v2.electric, v2.color)

	fmt.Println(v1.make, v1.model, v1.engine, v1.color)
	fmt.Println(v2.make, v2.model, v2.engine, v2.color)

	fmt.Println(v1.make, v1.model, v1.engine.electric, v1.color)
	fmt.Println(v2.make, v2.model, v2.engine.electric, v2.color)
}
