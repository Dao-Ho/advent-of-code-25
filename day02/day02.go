package day02

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input string) int {
	count := 0
	number_ranges := strings.Split(strings.TrimSpace(input), ",")
	for _, num_range := range number_ranges {
		start, end := parse_range(num_range)
		for i := start; i <= end; i++ {
			digits := numberToDigits(i)
			mid := len(digits) / 2
			left := digits[:mid]
			right := digits[mid:]
			if len(left) != len(right) {
				continue
			} else {
				if digitsToNumber(left) == digitsToNumber(right) {
					count += i
				}
			}
		}
	}

	return count
}

func parse_range(line string) (int, int) {
	parsed_range := strings.Split(strings.TrimSpace(line), "-")
	if len(parsed_range) != 2 {
		log.Fatal("Expected start and end value in {start}-{end}, got: ", line)
	}
	start, err1 := strconv.Atoi(parsed_range[0]) // Atoi returns (int, error)
	if err1 != nil {
		log.Fatal(err1)
	}
	end, err2 := strconv.Atoi(parsed_range[1]) // should be index [1], not [0]
	if err2 != nil {
		log.Fatal(err2)
	}
	return start, end
}

func numberToDigits(n int) []int {
	str := strconv.Itoa(n)
	digits := make([]int, len(str))
	for i, char := range str {
		digits[i] = int(char - '0')
	}
	return digits
}

func digitsToNumber(digits []int) int {
	str := ""
	for _, digit := range digits {
		str += strconv.Itoa(digit)
	}
	num, _ := strconv.Atoi(str)
	return num
}
