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
		name     string
		part     int
		fn       func([]byte) (string, error)
		expected string
	}{
		{"day1_sample", 1, day1part1, "3"},
		{"day1", 1, day1part1, "964"},
		{"day1_sample", 2, day1part2, "6"},
		{"day1", 2, day1part2, "5872"},

		{"day2_sample", 1, day2part1, "1227775554"},
		{"day2", 1, day2part1, "13919717792"},
		{"day2_sample", 2, day2part2, "4174379265"},
		{"day2", 2, day2part2, "14582313461"},

		{"day3_sample", 1, day3part1, "357"},
		{"day3", 1, day3part1, "17031"},
		{"day3_sample", 2, day3part2, "3121910778619"},
		{"day3", 2, day3part2, "168575096286051"},

		{"day4_sample", 1, day4part1, "13"},
		{"day4", 1, day4part1, "1489"},
		{"day4_sample", 2, day4part2, "43"},
		{"day4", 2, day4part2, "8890"},

		{"day5_sample", 1, day5part1, "3"},
		{"day5", 1, day5part1, "640"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s part %d", tt.name, tt.part), func(t *testing.T) {
			input, err := os.ReadFile("./inputs/" + tt.name + ".txt")
			require.NoError(t, err)
			result, err := tt.fn(input)
			require.NoError(t, err)
			assert.Equal(t, result, tt.expected)
		})
	}
}
