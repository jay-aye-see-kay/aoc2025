package aoc2025

import (
	"bytes"
	"fmt"
	"strconv"
)

type Point2D struct {
	X int
	Y int
}

func parseDay9(input []byte) ([]Point2D, error) {
	points := []Point2D{}

	for line := range bytes.Lines(input) {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		parts := bytes.Split(line, []byte{','})
		if len(parts) != 2 {
			return nil, fmt.Errorf("day 9 bad line")
		}
		x, err := strconv.Atoi(string(parts[0]))
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(string(parts[1]))
		if err != nil {
			return nil, err
		}
		points = append(points, Point2D{x, y})
	}

	return points, nil
}

func day9part1(input []byte) (int, error) {
	points, err := parseDay9(input)
	if err != nil {
		return 0, err
	}

	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			width := abs(points[j].X-points[i].X) + 1
			height := abs(points[j].Y-points[i].Y) + 1
			area := width * height
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
