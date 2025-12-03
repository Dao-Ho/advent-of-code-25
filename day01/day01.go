package day01

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input string) int {
	count := 0
	current_number := 50
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		direction := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Expected a turn amount as an integer, got ", line[1:])
		}
		switch direction {
		case 'R':
			current_number = (current_number + amount) % 100

		case 'L':
			subtracted_number := current_number - amount
			if subtracted_number < 0 {
				modulo_number := subtracted_number % 100
				current_number = 100 + modulo_number
				if current_number == 100 {
					current_number = 0
				}
			} else {
				current_number = subtracted_number
			}
		}
		if current_number == 0 {
			count += 1
		}
	}
	return count
}

func Part2(input string) int {
	count := 0
	current_number := 50
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		direction := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Expected a turn amount as an integer, got ", line[1:])
		}
		switch direction {
		case 'R':
			clicks := (current_number + amount) / 100

			count += clicks
			current_number = (current_number + amount) % 100

		case 'L':
			subtracted_number := current_number - amount
			if subtracted_number < 0 {
				clicks := (-subtracted_number / 100)
				if current_number != 0 {
					count += 1
				}
				count += clicks
				modulo_number := subtracted_number % 100
				current_number = 100 + modulo_number
				if current_number == 100 {
					current_number = 0
				}
			} else {
				current_number = subtracted_number
				if current_number == 0 {
					count += 1
				}

			}
		}
	}
	return count
}
