package aoc2025

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func day2part1(input []byte) (string, error) {
	isValid := func(id int) bool {
		id_str := strconv.Itoa(id)
		if len(id_str)%2 != 0 {
			return true // can't be a repeat if odd length
		}
		mid := len(id_str) / 2
		left, right := id_str[:mid], id_str[mid:]
		return left != right
	}

	return _day2(input, isValid)
}

func day2part2(input []byte) (string, error) {
	isValid := func(id int) bool {
		id_str := strconv.Itoa(id)
		chunk_groups := getEqualChunks(id_str)
		for _, chunks := range chunk_groups {
			if len(slices.Compact(chunks)) == 1 {
				return false
			}
		}
		return true
	}

	return _day2(input, isValid)
}

// split input string into all possible evenly sized chunks
func getEqualChunks(input string) [][]string {
	input_len := len(input)
	var factors []int
	for i := input_len - 1; i > 0; i-- {
		if input_len%i == 0 {
			factors = append(factors, i)
		}
	}

	var chunk_groups [][]string
	for _, factor := range factors {
		input_copy := input
		var chunks []string
		for len(input_copy) > 0 {
			chunks = append(chunks, input_copy[:factor])
			input_copy = input_copy[factor:]
		}
		chunk_groups = append(chunk_groups, chunks)
	}

	return chunk_groups
}

func _day2(input []byte, isValid func(int) bool) (string, error) {
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
