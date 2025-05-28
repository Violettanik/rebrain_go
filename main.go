package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer handlePanic()

	// Чтение данных из файла
	lines, err := readLines("data/in.txt")
	if err != nil {
		panic(fmt.Sprintf("error reading file: %v", err))
	}

	// Обработка данных и запись в файл
	outputFile, err := os.Create("data/data_out.txt")
	if err != nil {
		panic(fmt.Sprintf("error creating output file: %v", err))
	}
	defer outputFile.Close()

	for i, line := range lines {
		row := i + 1
		fields := strings.Split(line, "|")
		if len(fields) < 4 {
			panic(fmt.Sprintf("parse error: empty field on string %d", row))
		}

		for _, field := range fields {
			if strings.TrimSpace(field) == "" {
				panic(fmt.Sprintf("parse error: empty field on string %d", row))
			}
		}

		_, err := fmt.Fprintf(outputFile, "Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", 
			row, fields[0], fields[1], fields[2])
		if err != nil {
			panic(fmt.Sprintf("error writing to file: %v", err))
		}
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
		fmt.Println("\nContent of data_out.txt before panic:")
		content, err := os.ReadFile("data/data_out.txt")
		if err != nil {
			fmt.Println("Could not read output file:", err)
		} else {
			fmt.Println(string(content))
		}
	}
}
