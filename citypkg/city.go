package citypkg

import (
    "fmt"
    "/root/go/module03/01_task/wordz" // замените на ваш путь
)

func City() string {
    return wordz.Random()
}

func Digit() string {
    // Возвращает случайное число в виде строки: one, two, three...
    digits := []string{"one", "two", "three", "four", "five"}
    max := int64(len(digits))
    r, _ := rand.Int(rand.Reader, big.NewInt(max))
    return digits[r.Int64()]
}
