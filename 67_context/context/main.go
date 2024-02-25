package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context     : ", ctx)
	fmt.Println("Context Err : ", ctx.Err())
	fmt.Printf("Context Type : %T\n", ctx)
	fmt.Println("--------------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context     : ", ctx)
	fmt.Println("Context Err : ", ctx.Err())
	fmt.Printf("Context Type : %T\n", ctx)
	fmt.Println("cancel     : ", cancel)
	fmt.Printf("cancel type : %T\n", cancel)
	fmt.Println("--------------------")

	cancel()

	fmt.Println("Context     : ", ctx)
	fmt.Println("Context Err : ", ctx.Err()) // context canceled
	fmt.Printf("Context Type : %T\n", ctx)
	fmt.Println("cancel     : ", cancel)
	fmt.Printf("cancel type : %T\n", cancel)
	fmt.Println("--------------------")

}
