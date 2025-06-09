package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// LimitError - кастомная ошибка для превышения лимита строк
type LimitError struct {
	message     string
	limit       int
	lastString  string
	lineCount   int
}

// Error реализует интерфейс error
func (e *LimitError) Error() string {
	return fmt.Sprintf("%s, limit: %d, last string: %s", e.message, e.limit, e.lastString)
}

// countLines считает строки в файле с учетом лимита
func countLines(filePath string, limit int) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	var lastLine string

	for scanner.Scan() {
		lineCount++
		lastLine = scanner.Text()
		
		if lineCount > limit {
			return lineCount, &LimitError{
				message:    "string count exceed limit",
				limit:      limit,
				lastString: lastLine,
				lineCount:  lineCount,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		if errors.Is(err, io.EOF) {
			return lineCount, nil
		}
		return lineCount, fmt.Errorf("error reading file: %w", err)
	}

	return lineCount, nil
}

func main() {
	filePath := "data/in.txt"
	limit := 10 // Произвольное значение лимита

	count, err := countLines(filePath, limit)
	
	// Обработка ошибки превышения лимита
	var limitErr *LimitError
	if errors.As(err, &limitErr) {
		fmt.Printf("string count exceed limit, please read another file =) err: %v\n", err)
		return
	}
	
	// Обработка других ошибок
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Успешный случай
	fmt.Printf("Total strings: %d\n", count)
}
