package aoc2025

import (
	"fmt"
	"strconv"
	"strings"
)

func day1part1(input []byte) (string, error) {
	pointer := 50
	count_hit_zero := 0

	for line := range strings.Lines(string(input)) {
		line = strings.TrimSpace(line)
		direction := line[:1]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", err
		}

		switch direction {
		case "L":
			pointer -= count
		case "R":
			pointer += count
		default:
			return "", fmt.Errorf("unknown direction: %s", direction)
		}

		for pointer < 0 {
			pointer += 100
		}
		for pointer > 99 {
			pointer -= 100
		}

		if pointer == 0 {
			count_hit_zero += 1
		}
	}

	return strconv.Itoa(count_hit_zero), nil
}

func day1part2(input []byte) (string, error) {
	pointer := 50
	count_past_zero := 0

	for line := range strings.Lines(string(input)) {
		line = strings.TrimSpace(line)
		direction := line[:1]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", err
		}

		op := map[string]func(int) int{
			"L": func(i int) int { return i - 1 },
			"R": func(i int) int { return i + 1 },
		}[direction]

		for range count {
			pointer = op(pointer)
			switch pointer {
			case -1:
				pointer = 99
			case 100:
				pointer = 0
			}

			if pointer == 0 {
				count_past_zero += 1
			}
		}

	}

	return strconv.Itoa(count_past_zero), nil
}
