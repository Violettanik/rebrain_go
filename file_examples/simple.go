package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	const b = 10
	
	a = 5
	c := a + b
	
	fmt.Println(c)
	fmt.Println(math.Sqrt(float64(c)))
}
