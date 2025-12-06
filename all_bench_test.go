package aoc2025

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkAll(b *testing.B) {
	benchmarks := []struct {
		day  int
		part int
		fn   func([]byte) (int, error)
	}{
		{1, 1, day1part1},
		{1, 2, day1part2},
		{2, 1, day2part1},
		{2, 2, day2part2},
		{3, 1, day3part1},
		{3, 2, day3part2},
		{4, 1, day4part1},
		{4, 2, day4part2},
		{5, 1, day5part1},
		{5, 2, day5part2},
	}

	for _, bm := range benchmarks {
		name := fmt.Sprintf("day%d_part%d", bm.day, bm.part)
		b.Run(name, func(b *testing.B) {
			input, err := os.ReadFile(fmt.Sprintf("./inputs/day%d.txt", bm.day))
			if err != nil {
				b.Fatalf("failed to read input: %v", err)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := bm.fn(input)
				if err != nil {
					b.Fatalf("solution failed: %v", err)
				}
			}
		})
	}
}
