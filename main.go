package main

import "fmt"

func contains(a []string, x string) bool {
	for _, item := range a {
		if item == x {
			return true
		}
	}
	return false
}

func getMax(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	slice := []string{"apple", "banana", "cherry"}
	strToCheck := "banana"

	fmt.Printf("Содержится ли '%s' в слайсе? %v\n", strToCheck, contains(slice, strToCheck))

	maxNumber := getMax(1, 5, 3, 9, 2)
	fmt.Printf("Максимальное число: %d\n", maxNumber)
}
