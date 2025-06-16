package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Тесты для функции Pad
func TestPad(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		expected string
	}{
		{
			name:     "длина меньше исходной",
			input:    "Hello",
			length:   3,
			expected: "Hello",
		},
		{
			name:     "длина равна исходной",
			input:    "Hello",
			length:   5,
			expected: "Hello",
		},
		{
			name:     "короткое дополнение",
			input:    "Hi",
			length:   5,
			expected: "Hi Hello", // Предполагаем, что Pad добавляет " Hello" пока не достигнет длины
		},
		{
			name:     "длинное дополнение",
			input:    "X",
			length:   20,
			expected: "X Hello, world Hello", // Пример ожидаемого результата
		},
		{
			name:     "пустая строка",
			input:    "",
			length:   10,
			expected: " Hello, wor", // Обрезается до нужной длины
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Pad(tt.input, tt.length)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Бенчмарк для функции Pad
func BenchmarkPad(b *testing.B) {
	testCases := []struct {
		name   string
		length int
	}{
		{"short", 50},
		{"medium", 500},
		{"long", 5000},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Pad("test", tc.length)
			}
		})
	}
}

// Пример теста для другой функции (если есть)
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "madam", "madam"},
		{"normal string", "hello", "olleh"},
		{"unicode", "привет", "тевирп"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ReverseString(tt.input))
		})
	}
}
