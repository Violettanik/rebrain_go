package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Объявление переменных
	strValue := "104"
	intValue := 35

	// Преобразование строки в целое число
	intFromStr, err := strconv.Atoi(strValue)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
		return
	}

	// Преобразование целого числа в строку
	strFromInt := strconv.Itoa(intValue)

	fmt.Printf("Преобразованное значение из строки: %d\n", intFromStr)
	fmt.Printf("Преобразованное значение из целого числа: %s\n", strFromInt)
}
