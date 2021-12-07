package main

import (
	"aoc_2021/lib/helper"
	"math"
)

func part1(positions []int) (int, int) {
	min, max := helper.MinMax(positions)

	position := 0
	fuel := math.MaxInt64

	for i := min; i <= max; i++ {
		cost := 0
		for _, p := range positions {
			cost += helper.Abs(p - i)
		}
		if cost < fuel {
			fuel = cost
			position = i
		}
	}

	return position, fuel
}

func part2(positions []int) (int, int) {
	min, max := helper.MinMax(positions)

	position := 0
	fuel := math.MaxInt64

	for i := min; i <= max; i++ {
		cost := 0
		for _, p := range positions {
			for c := 1; c <= helper.Abs(p-i); c++ {
				cost += c
			}
		}
		if cost < fuel {
			fuel = cost
			position = i
		}
	}

	return position, fuel
}
