package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
)

func logTime(start time.Time) {
	duration := time.Since(start)
	fmt.Printf("Время выполнения программы: %v\n", duration)
}

func main() {
	start := time.Now()

	inputFile, err := os.Open("data/in.txt")
	if err != nil {
		fmt.Println("Ошибки при открытии файла:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Ошибки при создании файла:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	scanner := bufio.NewScanner(inputFile)

	lineCount := 0

	for scanner.Scan() {
		lineCount++
		_, err = writer.WriteString(fmt.Sprintf("%d: %s\n", lineCount, scanner.Text()))
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	writer.Flush()

	fileInfo, _ := outputFile.Stat()

	fmt.Printf("Записано строк: %d, байт: %d\n", lineCount, fileInfo.Size())

	logTime(start)
}
