package main

import (
	"fmt"
	"math"
)

type AmericanVelocity float64 // миль/ч
type EuropeanVelocity float64 // км/ч

func main() {
	var europeanVelocity EuropeanVelocity = convertToKmh(120.4) // м/с в км/ч
	var americanVelocity AmericanVelocity = convertToMph(130)    // м/с в милях/ч

	fmt.Printf("Скорость в км/ч: %.2f\n", europeanVelocity)
	fmt.Printf("Скорость в милях/ч: %.2f\n", americanVelocity)
}

func convertToKmh(mps float64) EuropeanVelocity {
	return EuropeanVelocity(mps * 3.6) // 1 м/с = 3.6 км/ч
}

func convertToMph(mps float64) AmericanVelocity {
	return AmericanVelocity(math.Round((mps * 2.23694)*100) / 100) // 1 м/с = 2.23694 миль/ч, округление до двух знаков после запятой 
}
