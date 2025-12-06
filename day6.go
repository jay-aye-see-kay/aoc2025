package aoc2025

import (
	"strconv"
	"strings"
)

type MathProb struct {
	nums []int
	sign string
}

func day6part1(input []byte) (int, error) {
	var numProbs int
	for line := range strings.Lines(string(input)) {
		numProbs = len(strings.Fields(line))
		break
	}

	probs := make([]*MathProb, numProbs)
	for line := range strings.Lines(string(input)) {
		for i, field := range strings.Fields(line) {
			if probs[i] == nil {
				probs[i] = &MathProb{}
			}
			if field == "+" || field == "*" {
				probs[i].sign = field
			} else {
				num, err := strconv.Atoi(field)
				if err != nil {
					return 0, err
				}
				probs[i].nums = append(probs[i].nums, num)
			}
		}
	}

	sum := 0
	for _, prob := range probs {
		switch prob.sign {
		case "+":
			ans := 0
			for _, num := range prob.nums {
				ans += num
			}
			sum += ans
		case "*":
			ans := 1
			for _, num := range prob.nums {
				ans *= num
			}
			sum += ans
		}
	}

	return sum, nil
}
