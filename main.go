package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "data/in.txt"

	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Построчное чтение файла
	for scanner.Scan() {
		lineCount++
	}

	// Проверка на ошибки (включая EOF)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Вывод количества строк
	fmt.Printf("Total strings: %d\n", lineCount)
}
