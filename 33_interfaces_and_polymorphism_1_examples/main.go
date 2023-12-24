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

func (s sqaure) getArea() float64 {
	var area float64
	area = float64(s.side * s.side)
	return area
}

type rectangle struct {
	length  int
	breadth int
}

func (r rectangle) getArea() float64 {
	var area float64
	area = float64(r.length * r.breadth)
	return area
}

type shape interface {
	getArea() float64
}

func getShapeArea(s shape) float64 {
	return s.getArea()
}

func main() {
	s1 := sqaure{
		side: 4,
	}

	r1 := rectangle{
		length:  4,
		breadth: 2,
	}

	fmt.Println("s1.getArea() ", s1.getArea())
	fmt.Println("getShapeArea(s1) ", getShapeArea(s1))
	fmt.Println("getShapeArea(r1) ", getShapeArea(r1))

}
