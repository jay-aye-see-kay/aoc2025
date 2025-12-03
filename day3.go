package aoc2025

import (
	"strconv"
	"strings"
)

func day3part1(input []byte) (string, error) {
	sum := 0
	for line := range strings.Lines(string(input)) {
		nums, err := splitLine(strings.TrimSpace(line))
		if err != nil {
			return "", err
		}

		max1 := 0
		max1Index := 0
		for i, num := range nums {
			isLast := i >= len(nums)-1
			if !isLast && max1 < num {
				max1 = num
				max1Index = i
			}
		}

		max2 := 0
		for _, num := range nums[max1Index+1:] {
			if max2 < num {
				max2 = num
			}
		}

		result_str := strconv.Itoa(max1) + strconv.Itoa(max2)
		result, err := strconv.Atoi(result_str)
		if err != nil {
			return "", nil
		}
		sum += result

	}
	return strconv.Itoa(sum), nil
}

func splitLine(line string) ([]int, error) {
	nums := []int{}
	for _, char := range []rune(line) {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}
