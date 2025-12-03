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
	got := Part2(string(example))
	want := 6
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func TestPart2EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "single rotation to zero",
			input:    "L50",
			expected: 1, // ends at 0
		},
		{
			name:     "single rotation not to zero",
			input:    "R10",
			expected: 0, // 50 -> 60, never passes 0
		},
		{
			name:     "rotation passing through zero once",
			input:    "L60",
			expected: 1, // 50 -> ... -> 0 -> ... -> 90
		},
		{
			name:     "rotation passing through zero twice",
			input:    "L150",
			expected: 2, // 50 -> ... -> 0 -> ... -> 0 -> ... -> 0 (wraps 1.5 times)
		},
		{
			name:     "wrap around from 0 to 99 passing through 0",
			input:    "L50\nL1",
			expected: 1, // 50 -> 0 (count 1), then 0 -> 99 (should not count since we're already on 0)
		},
		{
			name:     "wrap around from 99 to 0",
			input:    "R49\nR1",
			expected: 1, // 50 -> 99, then 99 -> 0
		},
		{
			name:     "large rotation wrapping multiple times",
			input:    "R250",
			expected: 3, // 50 + 250 = 300, passes 0 at 100, 200, 300
		},
		{
			name:     "exactly 100 clicks (full circle)",
			input:    "L100",
			expected: 1, // 50 -> ... -> 0 -> ... -> 50
		},
		{
			name:     "rotation of zero",
			input:    "R0",
			expected: 0, // stays at 50
		},
		{
			name:     "multiple rotations with clicks through zero",
			input:    "L68\nL30\nR48",
			expected: 2, // L68: 50->82 passes 0 once; L30: 82->52; R48: 52->0 ends at 0
		},
		{
			name:     "never passes through zero",
			input:    "R10\nR10\nR10",
			expected: 0, // 50 -> 60 -> 70 -> 80
		},
		{
			name:     "rotation ending at zero counts once",
			input:    "R50",
			expected: 1, // 50 -> 100 (wraps to 0), counts once
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)
			if got != tt.expected {
				t.Errorf("Part2() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	result := Part2(string(input))
	t.Logf("Part 2: %d", result)
}


