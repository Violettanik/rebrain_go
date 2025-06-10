package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	n_44 := 44
        n_45 := 45
        go print_fib(n_44)
         print_fib(n_45)
	//fmt.Printf("\rFibonacci(%d) = %d\n_1", n_1, fib_44, "\rFibonacci(%d) = %d\n_2", n_2, fib_45)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func print_fib(x int) {
        fib_x := fib(x)
        fmt.Printf("\rFibonacci(%d) = %d\n", x, fib_x)
}
