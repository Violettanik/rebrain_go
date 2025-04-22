package main

import "fmt"

func main() {
	var A *int
	B := 10

	A = &B
	fmt.Println(*A)

	*A = 20
	fmt.Println(B)
}
