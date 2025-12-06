package aoc2025

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_all(t *testing.T) {
	tests := []struct {
		day      int
		part     int
		sample   bool
		fn       func([]byte) (string, error)
		expected string
	}{
		{1, 1, true, day1part1, "3"},
		{1, 1, false, day1part1, "964"},
		{1, 2, true, day1part2, "6"},
		{1, 2, false, day1part2, "5872"},

		{2, 1, true, day2part1, "1227775554"},
		{2, 1, false, day2part1, "13919717792"},
		{2, 2, true, day2part2, "4174379265"},
		{2, 2, false, day2part2, "14582313461"},

		{3, 1, true, day3part1, "357"},
		{3, 1, false, day3part1, "17031"},
		{3, 2, true, day3part2, "3121910778619"},
		{3, 2, false, day3part2, "168575096286051"},

		{4, 1, true, day4part1, "13"},
		{4, 1, false, day4part1, "1489"},
		{4, 2, true, day4part2, "43"},
		{4, 2, false, day4part2, "8890"},

		{5, 1, true, day5part1, "3"},
		{5, 1, false, day5part1, "640"},
		{5, 2, true, day5part2, "14"},
		{5, 2, false, day5part2, "365804144481581"},
	}

	for _, tt := range tests {
		filename := fmt.Sprintf("day%d", tt.day)
		testName := filename
		if tt.sample {
			filename += "_sample"
			testName += " sample"
		}
		testName = fmt.Sprintf("%s part %d", testName, tt.part)

		t.Run(testName, func(t *testing.T) {
			input, err := os.ReadFile("./inputs/" + filename + ".txt")
			require.NoError(t, err)
			result, err := tt.fn(input)
			require.NoError(t, err)
			assert.Equal(t, result, tt.expected)
		})
	}
}
