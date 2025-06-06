package citypkg

import (
    "gopackages-layout/internal/wordz"
    "strings"
)

var cities = []string{
    "Moscow",
    "London",
    "Paris",
    "Tokyo",
    "Berlin",
}

func City() string {
    randomStr := wordz.Random()
    word := strings.TrimPrefix(randomStr, wordz.Prefix)
    
    // Находим индекс слова в wordz.Words
    var index int
    for i, w := range wordz.Words {
        if w == word {
            index = i
            break
        }
    }
    
    // Возвращаем город по тому же индексу
    return cities[index]
}

func Digit() string {
    randomStr := wordz.Random()
    word := strings.TrimPrefix(randomStr, wordz.Prefix)
    return strings.ToLower(word)
}
