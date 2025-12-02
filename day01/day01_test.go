package day01

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	example, _ := os.ReadFile("example.txt")
	got := Part1(string(example))
	want := 3
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	result := Part1(string(input))
	t.Logf("Part 1: %d", result)
}

func TestPart1EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "single rotation to zero",
			input:    "L50",
			expected: 1,
		},
		{
			name:     "single rotation not to zero",
			input:    "R10",
			expected: 0,
		},
		{
			name:     "multiple zeros",
			input:    "L50\nR100\nL100",
			expected: 3, // hits 0 three times
		},
		{
			name:     "wrap around from 0 to 99",
			input:    "L50\nL1",
			expected: 1, // 50 -> 0 -> 99
		},
		{
			name:     "wrap around from 99 to 0",
			input:    "R49\nR1",
			expected: 1, // 50 -> 99 -> 0
		},
		{
			name:     "large rotation wrapping multiple times",
			input:    "R250", // wraps 2.5 times
			expected: 1,      // 50 + 250 = 300 % 100 = 0
		},
		{
			name:     "exactly 100 clicks (full circle)",
			input:    "L100",
			expected: 0, // stays at 50
		},
		{
			name:     "rotation of zero",
			input:    "R0",
			expected: 0, // stays at 50
		},
		{
			name:     "alternating left and right hitting zero",
			input:    "L50\nR50\nL50",
			expected: 2, // 50 -> 0 -> 50 -> 0
		},
		{
			name:     "never hits zero",
			input:    "R10\nR10\nR10",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
			if got != tt.expected {
				t.Errorf("Part1() = %v, want %v", got, tt.expected)
			}
		})
	}
}


func TestPart2Example(t *testing.T) {
	example, _ := os.ReadFile("example.txt")
	got := Part1(string(example))
	want := 6
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}