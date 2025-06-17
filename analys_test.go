package analys

import (
	"testing"
)

func TestAnalys(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected AnalysResult
	}{
		{
			name:     "simple file",
			filePath: "file_examples/simple.go",
			expected: AnalysResult{
				DeclCount:    2,
				CallCount:    4,
				AssignCount:  1,
				ImportsCount: 2,
			},
		},
		{
			name:     "complex file",
			filePath: "file_examples/complex.go",
			expected: AnalysResult{
				DeclCount:    6,
				CallCount:    16,
				AssignCount:  4,
				ImportsCount: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Analys(tt.filePath)
			if err != nil {
				t.Fatalf("Analys() error = %v", err)
			}

			if result != tt.expected {
				t.Errorf("Analys() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}
