package day02

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	example, _ := os.ReadFile("example.txt")
	got := Part1(string(example))
	want := 1227775554
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

// func TestPart1(t *testing.T) {
// 	input, _ := os.ReadFile("input.txt")
// 	result := Part1(string(input))
// 	t.Logf("Part 1: %d", result)
// }