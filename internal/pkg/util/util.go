package util
import "strings"
func Pad(s string, length int) string {
    if len(s) >= length {
        return s
    }
    
    var builder strings.Builder
    builder.WriteString(s)
    
    padStr := " Hello, world"
    for builder.Len() < length {
        builder.WriteString(padStr)
    }
    
    // Обрезаем если превысили длину
    if builder.Len() > length {
        return builder.String()[:length]
    }
    return builder.String()
}
