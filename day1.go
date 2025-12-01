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
