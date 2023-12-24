package main

import "fmt"

func variadicParams(nums ...int) int {
	fmt.Println(nums)
	fmt.Printf("%T\n",nums)
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	sum1 := variadicParams(4, 6, 2, 3, 3)
	fmt.Println(sum1)

	sum2 := variadicParams()
	fmt.Println(sum2)
}
