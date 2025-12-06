package aoc2025

import (
	"slices"
	"strconv"
	"strings"
)

func day5parse(input []byte) ([]Range, []int, error) {
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

func day5part1(input []byte) (int, error) {
	freshRanges, ingredientIDs, err := day5parse(input)
	if err != nil {
		return 0, err
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

	return count, nil
}

func day5part2(input []byte) (int, error) {
	freshRanges, _, err := day5parse(input)
	if err != nil {
		return 0, err
	}

	moreToMerge := true
	for moreToMerge {
		freshRanges, moreToMerge = mergeOverlapping(freshRanges)
	}

	count := 0
	for _, r := range freshRanges {
		count += r.count()
	}

	return count, nil
}

func mergeOverlapping(ranges []Range) ([]Range, bool) {
	for i, r1 := range ranges {
		for _, r2 := range ranges[i+1:] {
			if r1.overlaps(&r2) {
				// if we've found an overlap, delete the two old ones, add in the merged and return
				// starting again after each merge simplifies iterating array indexes
				newRanges := slices.DeleteFunc(ranges, func(r Range) bool {
					return r.is(r1) || r.is(r2)
				})
				newRanges = append(newRanges, r1.merge(&r2))
				return newRanges, true
			}
		}
	}
	return ranges, false
}

type Range struct {
	From int
	To   int
}

func (r *Range) contains(n int) bool {
	return n >= r.From && n <= r.To
}

func (r *Range) overlaps(r2 *Range) bool {
	overlaps := max(r.From, r2.From) <= min(r.To, r2.To)
	return overlaps
}

func (r *Range) merge(r2 *Range) Range {
	return Range{
		From: min(r.From, r2.From),
		To:   max(r.To, r2.To),
	}
}

func (r *Range) is(r2 Range) bool {
	return r.From == r2.From && r.To == r2.To
}

func (r *Range) count() int {
	return r.To - r.From + 1
}
