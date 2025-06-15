package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
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

// TestContainsDuplicate - table-driven тесты
func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: false,
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "has duplicates",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "multiple duplicates",
			input:    []int{1, 1, 2, 2},
			expected: true,
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsDuplicate(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsPalindrome - closure-driven тесты
func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		testFunc func(*testing.T)
	}{
		{
			name: "positive palindrome",
			testFunc: func(t *testing.T) {
				assert.True(t, IsPalindrome(121))
			},
		},
		{
			name: "negative number",
			testFunc: func(t *testing.T) {
				assert.False(t, IsPalindrome(-121))
			},
		},
		{
			name: "non-palindrome",
			testFunc: func(t *testing.T) {
				assert.False(t, IsPalindrome(123))
			},
		},
		{
			name: "single digit",
			testFunc: func(t *testing.T) {
				assert.True(t, IsPalindrome(9))
			},
		},
		{
			name: "zero",
			testFunc: func(t *testing.T) {
				assert.True(t, IsPalindrome(0))
			},
		},
		{
			name: "large palindrome",
			testFunc: func(t *testing.T) {
				assert.True(t, IsPalindrome(123454321))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testFunc)
	}
}
func TestReverseInt_UnsupportedType(t *testing.T) {
    // Используем анонимную функцию для отлова panic
    assert.PanicsWithValue(t, "unsupported type", func() {
        ReverseInt("123") // Передаем строку вместо int
    }, "Функция должна вызывать panic при неподдерживаемом типе")

    // Можно добавить другие неподдерживаемые типы
    assert.PanicsWithValue(t, "unsupported type", func() {
        ReverseInt(12.34) // float64
    })

    assert.PanicsWithValue(t, "unsupported type", func() {
        ReverseInt([]int{1, 2, 3}) // slice
    })
}
