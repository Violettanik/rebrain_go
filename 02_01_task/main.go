package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2) // 2 потоков (P)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Starting goroutine %d \n", i+1)
			for {} // Бесконечный цикл (нагружает CPU)
		}(i)
	}
	select {} // Блокируем main()
}
