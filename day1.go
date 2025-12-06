package aoc2025

import (
	"strconv"
	"strings"
)

func day1part1(input []byte) (int, error) {
	zero_stops, _, err := _day1(input)
	if err != nil {
		return 0, err
	}
	return zero_stops, nil

}

func day1part2(input []byte) (int, error) {
	_, zero_passes, err := _day1(input)
	if err != nil {
		return 0, err
	}
	return zero_passes, nil
}

func _day1(input []byte) (int, int, error) {
	pointer := 50
	zero_stops := 0
	zero_passes := 0

	for line := range strings.Lines(string(input)) {
		line = strings.TrimSpace(line)
		direction := line[:1]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, 0, err
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
				zero_passes += 1
			}
		}

		if pointer == 0 {
			zero_stops += 1
		}
	}

	return zero_stops, zero_passes, nil
}
