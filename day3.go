package aoc2025

import (
	"strconv"
	"strings"
)

func day3part1(input []byte) (int, error) {
	return _day3(input, 2)
}

func day3part2(input []byte) (int, error) {
	return _day3(input, 12)
}

func _day3(input []byte, goalLen int) (int, error) {
	sum := 0
	for line := range strings.Lines(string(input)) {
		nums, err := splitNumAsStr(line)
		if err != nil {
			return 0, err
		}

		maxNums := []int{}
		start := 0
		for n := goalLen; n > 0; n-- {
			end := len(nums) - n
			foundIndex := getNthMaxIndex(nums, start, end)
			maxNums = append(maxNums, nums[foundIndex])
			start = foundIndex + 1
		}

		result, err := joinNumsAsStr(maxNums)
		if err != nil {
			return 0, err
		}
		sum += result
	}
	return sum, nil
}

func getNthMaxIndex(nums []int, start, end int) int {
	maxIndex := start
	for i := start; i <= end; i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}
	return maxIndex
}

func splitNumAsStr(line string) ([]int, error) {
	nums := []int{}
	for char := range strings.SplitSeq(strings.TrimSpace(line), "") {
		num, err := strconv.Atoi(char)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func joinNumsAsStr(nums []int) (int, error) {
	s := ""
	for _, num := range nums {
		s += strconv.Itoa(num)
	}
	return strconv.Atoi(s)

}
