package main

import (
	"fmt"
	"math"
)

func main() {
	circumference := 35.0
	radius := circumference / (2 * math.Pi)

	R := &radius

	area := math.Pi * (*R) * (*R)

	fmt.Printf("Площадь круга с радиусом %.2f равна %.2f\n", *R, area)
}
