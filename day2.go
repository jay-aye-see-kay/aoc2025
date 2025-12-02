package aoc2025

import (
	"fmt"
	"strconv"
	"strings"
)

func day2part1(input []byte) (string, error) {
	sum := 0

	for pair := range strings.SplitSeq(string(input), ",") {
		split_pairs := strings.Split(strings.TrimSpace(pair), "-")
		if len(split_pairs) != 2 {
			return "", fmt.Errorf("invalid pair %s", pair)
		}
		start_str, end_str := split_pairs[0], split_pairs[1]
		start, err := strconv.Atoi(start_str)
		if err != nil {
			return "", err
		}
		end, err := strconv.Atoi(end_str)
		if err != nil {
			return "", err
		}
		for i := start; i <= end; i++ {
			if !isValid(i) {
				sum += i
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func isValid(id int) bool {
	id_str := strconv.Itoa(id)
	if len(id_str)%2 != 0 {
		return true // can be a repeat if odd length
	}
	mid := len(id_str) / 2
	left, right := id_str[:mid], id_str[mid:]
	return left != right
}
