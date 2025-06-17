package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	pi       = 3.14
	appName  = "AST Analyzer"
)

var (
	counter int
	results []string
)

type Config struct {
	Timeout int
	Retries int
}

func calculate(x, y float64) float64 {
	return math.Pow(x, y) + math.Sqrt(x*y)
}

func process(input string) string {
	parts := strings.Split(input, ",")
	var output strings.Builder

	for i, part := range parts {
		if i > 0 {
			output.WriteString("|")
		}
		output.WriteString(strings.TrimSpace(part))
	}

	return output.String()
}

func main() {
	counter = 10
	results = make([]string, 0, 5)

	cfg := Config{
		Timeout: 30,
		Retries: 3,
	}

	fmt.Println("Starting", appName)
	fmt.Printf("Config: %+v\n", cfg)

	res1 := calculate(2.5, 3.0)
	res2 := process("a, b, c , d")

	results = append(results, fmt.Sprintf("res1: %.2f", res1))
	results = append(results, "res2: "+res2)

	for _, r := range results {
		fmt.Println(r)
	}
}
