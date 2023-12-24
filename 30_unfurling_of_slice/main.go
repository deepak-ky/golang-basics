package main

import "fmt"

//func add(constant int, nums []int) int { // would not work  -> cannot use ... in call to non-variadic add
//func add(constant int, x,y,z int) int { // would not work
func add(constant int, nums ...int) int {
	// add stuff to constant here, would not change its value in main
	for _, num := range nums {
		constant += num
	}
	return constant
}

//function def nums ...int -> then function call should be numbers...
//function def nums []int -> then function call should be numbers
//criss-cross would not work

func main() {

	constant := 10
	numbers := []int{2,3,4}
	sum := add(constant, numbers...) // unfurling of a slice
	fmt.Println(constant)
	fmt.Println(sum)



}
