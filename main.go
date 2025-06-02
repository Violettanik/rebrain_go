package main

import (
    "fmt"
//    "log"

    "01_task/citypkg"
//    "01_task/wordz"
    "github.com/huandu/xstrings" // замените на актуальный пакет
)

func main() {
    city := citypkg.City()
    digit := citypkg.Digit()

    fmt.Println("Original city:", city)
    
    // Предположим, что функция Shuffle есть в пакете xstrings
    shuffledCity := xstrings.Shuffle(city)
    
    fmt.Println("Shuffled city:", shuffledCity)
    
    fmt.Println("Digit:", digit)
}
