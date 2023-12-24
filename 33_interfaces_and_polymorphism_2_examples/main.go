package main

import "fmt"

/*
For values of struct types, to use them as interface type
    a. the struct should implement all interface functions
    b. all functions
        i.   should have same name
        ii.  same number of parameters
        iii. same type of parameters
        iv.  same return type
*/

type sqaure struct {
	side int
}

func (s sqaure) multiplySideByX(x int) float32 {
	var ans float32
	ans = float32(s.side * x)
	return ans
}

func (s sqaure) getLargestSide() int {
	return s.side
}

type shape interface {
	multiplySideByX(int) float32
	//getSmallestSide() int  -> If uncommented and square does not implement this func, square type values cannot be of type shape
}

func multiplyShapeSideByX(s shape, x int) float32 {
	return s.multiplySideByX(x)
}

func main() {
	s1 := sqaure{
		side: 4,
	}

	fmt.Println("s1:", s1)
	fmt.Println("s1.multiplySideByX(3)", s1.multiplySideByX(3))
	fmt.Println("s1.getLargestSide()", s1.getLargestSide())

	fmt.Println("multiplyShapeSideByX(s1,3)", multiplyShapeSideByX(s1, 3))

}
