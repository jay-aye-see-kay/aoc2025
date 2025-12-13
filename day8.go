package aoc2025

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
)

type Point struct {
	X int
	Y int
	Z int
}

type Pair struct {
	p1          *Point
	p2          *Point
	distSquared int
}

type Circuit struct {
	points map[*Point]struct{}
}

type Circuits []*Circuit

func day8part1(input []byte) (int, error) {
	points, err := day8Parse(input)
	if err != nil {
		return 0, err
	}
	// sample input length is 230, real is 17696, use that to detect how many pair to join into circuits
	pairs := 1000
	if len(input) < 1000 {
		pairs = 10
	}
	shortestPairs := sortedPairs(points, pairs)

	var circuits Circuits
	for _, pair := range shortestPairs {
		addToCircuits(pair, &circuits, nil)
	}

	slices.SortFunc(circuits, func(c1 *Circuit, c2 *Circuit) int {
		return len(c2.points) - len(c1.points)
	})

	answer := 1
	for _, c := range circuits[:3] {
		answer *= len(c.points)
	}
	return answer, nil
}

func day8part2(input []byte) (int, error) {
	points, err := day8Parse(input)
	if err != nil {
		return 0, err
	}
	pairs := sortedPairs(points, -1)

	pointSet := make(map[*Point]struct{})
	var circuits Circuits

	for _, pair := range pairs {
		addToCircuits(pair, &circuits, pointSet)
		if len(pointSet) == len(points) {
			return pair.p1.X * pair.p2.X, nil
		}
	}

	return 0, fmt.Errorf("did not find last pair to complete one circuit")
}

func day8Parse(input []byte) ([]Point, error) {
	var points []Point
	for line := range bytes.Lines(input) {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		strParts := bytes.Split(line, []byte{','})
		if len(strParts) != 3 {
			return nil, fmt.Errorf("day 8 bad line")
		}
		intParts := make([]int, 3)
		for i, strPart := range strParts {
			intPart, err := strconv.Atoi(string(strPart)) // OPTIMIZE ME
			if err != nil {
				return nil, err
			}
			intParts[i] = intPart
		}
		points = append(points, Point{
			X: intParts[0],
			Y: intParts[1],
			Z: intParts[2],
		})

	}
	return points, nil
}

// IDEA: use heap with a limit so we don't have to sort 1mil pairs
// https://pkg.go.dev/container/heap#example-package-IntHeap
func sortedPairs(points []Point, limit int) []*Pair {
	var pairs []*Pair
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := &points[i]
			p2 := &points[j]
			distSquared := square(p1.X-p2.X) + square(p1.Y-p2.Y) + square(p1.Z-p2.Z)
			pairs = append(pairs, &Pair{p1, p2, distSquared})
		}
	}
	slices.SortFunc(pairs, func(p1 *Pair, p2 *Pair) int {
		return p1.distSquared - p2.distSquared
	})
	if limit < 0 {
		return pairs
	}
	return pairs[:limit]
}

func addToCircuits(pair *Pair, circuits *Circuits, pointSet map[*Point]struct{}) {
	if pointSet != nil {
		pointSet[pair.p1] = struct{}{}
		pointSet[pair.p2] = struct{}{}
	}

	var p1FoundAt *Circuit
	var p2FoundAt *Circuit

	for _, circuit := range *circuits {
		_, p1Found := circuit.points[pair.p1]
		_, p2Found := circuit.points[pair.p2]

		// CASE 1: both connected in this circuit, do nothing
		if p1Found && p2Found {
			return
		}

		// CASE 2: one in this circuit, add the other point, and the pair
		if p1Found {
			circuit.points[pair.p2] = struct{}{}
			p1FoundAt = circuit
		}
		if p2Found {
			circuit.points[pair.p1] = struct{}{}
			p2FoundAt = circuit
		}
	}

	// CASE 3: found in two circuts, merge!
	if p1FoundAt != nil && p2FoundAt != nil {
		// merge second circuit into first
		for k := range p2FoundAt.points {
			p1FoundAt.points[k] = struct{}{}
		}
		// delete second circuit
		*circuits = slices.DeleteFunc(*circuits, func(c *Circuit) bool {
			return c == p2FoundAt
		})
	}

	// CASE 2 (part 2), found one, added it, can exit
	if p1FoundAt != nil || p2FoundAt != nil {
		return
	}

	// CASE 4: not in any circuit create new one
	*circuits = append(*circuits, &Circuit{
		points: map[*Point]struct{}{
			pair.p1: {},
			pair.p2: {},
		},
	})
}

func square(x int) int {
	return x * x
}
