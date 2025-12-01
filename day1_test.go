package aoc2025

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var sample_input = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func Test_day1part1(t *testing.T) {
	t.Run("sample data", func(t *testing.T) {
		input, err := os.ReadFile("./inputs/day1_sample.txt")
		require.NoError(t, err)
		assert.Equal(t, day1part1(input), 3)
	})
}
