package aoc2025

import (
	"strconv"
	"strings"
)

func day5part1(input []byte) (string, error) {
	freshRanges, ingredientIDs, err := parse(input)
	if err != nil {
		return "", err
	}

	count := 0
	for _, id := range ingredientIDs {
		for _, r := range freshRanges {
			if r.contains(id) {
				count += 1
				break
			}
		}
	}

	return strconv.Itoa(count), nil
}

func parse(input []byte) ([]Range, []int, error) {
	p := strings.Split(string(input), "\n\n")
	rangesStr, idsStr := p[0], p[1]

	var ranges []Range
	for line := range strings.Lines(rangesStr) {
		p := strings.Split(strings.TrimSpace(line), "-")
		from, err := strconv.Atoi(p[0])
		if err != nil {
			return nil, nil, err
		}
		to, err := strconv.Atoi(p[1])
		if err != nil {
			return nil, nil, err
		}
		ranges = append(ranges, Range{from, to})
	}

	var ids []int
	for line := range strings.Lines(idsStr) {
		id, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, nil, err
		}
		ids = append(ids, id)
	}

	return ranges, ids, nil
}

type Range struct {
	From int
	To   int
}

func (r *Range) contains(n int) bool {
	return n >= r.From && n <= r.To
}
