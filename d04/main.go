package main

import (
	"aoc_2021/lib/helper"
	"aoc_2021/lib/set"
	"strings"
)

type Board struct {
	data [5][5]int
}

func (b *Board) Check(drawn *set.IntSet) bool {
	// Check Rows
	for x := 0; x < 5; x++ {
		valid := true
		for y := 0; y < 5; y++ {
			if !drawn.Has(b.data[x][y]) {
				valid = false
				break
			}
		}
		if valid {
			return true
		}
	}

	// Check Columns
	for y := 0; y < 5; y++ {
		valid := true
		for x := 0; x < 5; x++ {
			if !drawn.Has(b.data[x][y]) {
				valid = false
				break
			}
		}
		if valid {
			return true
		}
	}
	return false
}

func (b *Board) Sum(drawn *set.IntSet) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !drawn.Has(b.data[x][y]) {
				sum += b.data[x][y]
			}
		}
	}
	return sum
}

func ParseData(lines []string) ([]int, []*Board) {
	drawn := make([]int, 0)
	for _, v := range strings.Split(lines[0], ",") {
		drawn = append(drawn, helper.Atoi(v))
	}
	boards := make([]*Board, 0)
	for i := 2; i < len(lines); i += 6 {
		b := parseBoard(lines[i : i+5])
		boards = append(boards, b)
	}
	return drawn, boards
}

func parseBoard(lines []string) *Board {
	data := [5][5]int{}
	for i := 0; i < 5; i++ {
		for j, v := range strings.Split(strings.ReplaceAll(strings.TrimSpace(lines[i]), "  ", " "), " ") { // this is so bad
			d := helper.Atoi(v)
			data[i][j] = d
		}
	}
	return &Board{
		data: data,
	}
}

func part1(drawn []int, boards []*Board) int {
	drawnSet := set.NewIntSet()
	for _, i := range drawn {
		drawnSet.Add(i)
		for _, b := range boards {
			if b.Check(drawnSet) {
				return b.Sum(drawnSet) * i
			}
		}
	}
	return 0
}

func part2(drawn []int, boards []*Board) int {
	drawnSet := set.NewIntSet()
	hasWon := set.NewIntSet()
	for _, i := range drawn {
		drawnSet.Add(i)
		for bIdx, b := range boards {
			if b.Check(drawnSet) {
				if !hasWon.Has(bIdx) && hasWon.Size() == len(boards)-1 {
					// Last to win
					return b.Sum(drawnSet) * i
				}
				hasWon.Add(bIdx)
			}
		}
	}
	return 0
}
