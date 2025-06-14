package main

func ReverseInt(x interface{}) int {
    var num int
    switch v := x.(type) {
    case int:
        num = v
    default:
        panic("unsupported type")
    }

    negative := false
    if num < 0 {
        negative = true
        num = -num
    }

    rev := 0
    for num > 0 {
        rev = rev*10 + num%10
        num /= 10
    }

    if negative {
        return -rev
    }
    return rev
}

// ContainsDuplicate проверяет наличие дубликатов в слайсе
func ContainsDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}

// IsPalindrome проверяет, является ли число палиндромом
func IsPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    original := x
    reversed := 0
    for x > 0 {
        reversed = reversed*10 + x%10
        x /= 10
    }
    return original == reversed
}
