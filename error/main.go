package main

import (
	"fmt"
)

type CustomError struct {
	errMessage string
	errCode    int
}

func (ce CustomError) Error() string {
	return ce.errMessage
}

func divide(n1, n2 int) (float64, error) {
	var ans float64
	if n2 == 0 {
		return ans, CustomError{
			errMessage: "trying to divide by zero",
			errCode:    123,
		}
	}

	return float64(n1 / n2), nil
}

func main() {

	var emptyCustomError CustomError
	fmt.Println("emptyCustomError", emptyCustomError)

	ans, err := divide(4, 0)
	fmt.Println("ans", ans)
	fmt.Println("customError", err)
	
	if err != nil {
		cErr := err.(CustomError)
		fmt.Println("Could not divide due to :", cErr.errMessage)
		fmt.Println("Error code is :",cErr.errCode)
	} else {
		fmt.Println(ans)
	}
}
