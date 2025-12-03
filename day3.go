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

func day3part2(input []byte) (string, error) {
	sum := 0
	for line := range strings.Lines(string(input)) {
		nums, err := splitLine(strings.TrimSpace(line))
		if err != nil {
			return "", err
		}

		goalLen := 12
		maxNums := []int{}
		start := 0
		for n := goalLen; n > 0; n-- {
			end := len(nums) - n
			newMaxIndex := getNthMax(nums, start, end)
			maxNums = append(maxNums, nums[newMaxIndex])
			start = newMaxIndex + 1
		}

		result, err := joinNumAsStr(maxNums)
		if err != nil {
			return "", err
		}
		sum += result
	}
	return strconv.Itoa(sum), nil
}

func getNthMax(nums []int, start, end int) int {
	maxIndex := start
	for i := start; i <= end; i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}
	return maxIndex
}

func joinNumAsStr(nums []int) (int, error) {
	s := ""
	for _, num := range nums {
		s += strconv.Itoa(num)
	}
	return strconv.Atoi(s)

}
