package main

import (
	"fmt"
	"gopackages-layout/internal/citypkg"
	"github.com/huandu/xstrings"
)

func main() {
	city := citypkg.City()
	digit := citypkg.Digit()

	fmt.Println("Original city:", city)
	
	shuffledCity := xstrings.Shuffle(city)
	
	fmt.Println("Shuffled city:", shuffledCity)
	
	fmt.Println("Digit:", digit)
}
