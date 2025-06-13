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
