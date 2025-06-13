package main

import (
    "testing"
//    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
    // Тест-кейсы с использованием testify require
    t.Run("Positive number", func(t *testing.T) {
        result := ReverseInt(123)
        require.Equal(t, 321, result)
    })

    t.Run("Negative number", func(t *testing.T) {
        result := ReverseInt(-649)
        require.Equal(t, -946, result)
    })

    t.Run("Zero", func(t *testing.T) {
        result := ReverseInt(0)
        require.Equal(t, 0, result)
    })

    // Дополнительные тесты
    t.Run("Number with trailing zeros", func(t *testing.T) {
        result := ReverseInt(1200)
        require.Equal(t, 21, result) // предполагается, что ведущие нули после реверса игнорируются
    })

    t.Run("Large number", func(t *testing.T) {
        result := ReverseInt(1000000001)
        require.Equal(t, 1000000001, result) // реверс совпадает с исходным
    })

    t.Run("Negative large number", func(t *testing.T) {
        result := ReverseInt(-987654321)
        require.Equal(t, -123456789, result)
    })

    t.Run("Single digit", func(t *testing.T) {
        result := ReverseInt(7)
        require.Equal(t, 7, result)
    })
}
