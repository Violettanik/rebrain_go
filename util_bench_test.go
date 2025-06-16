package util

import "testing"// Оптимизированные версии
func FibIterative(n int) int {
    a, b := 0, 1
    for i := 0; i < n; i++ {
        a, b = b, a+b
    }
    return a
}

func MakeSliceOptimized(l int) []int {
    s := make([]int, 0, l)
    for i := 0; i < l; i++ {
        s = append(s, i)
    }
    return s
}

// Бенчмарки для оптимизированных версий
func BenchmarkFibIterative(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FibIterative(35)
    }
}

func BenchmarkMakeSliceOptimized(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MakeSliceOptimized(3000000)
    }
}
