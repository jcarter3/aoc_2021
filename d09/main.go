package main

import (
	"aoc_2021/lib/grid"
	"aoc_2021/lib/helper"
	"sort"
	"strings"
)

func ParseGrid(lines []string) [][]int {
	g := make([][]int, len(lines), len(lines))
	for i, l := range lines {
		row := make([]int, len(l), len(l))
		for j, s := range strings.Split(l, "") {
			row[j] = helper.Atoi(s)
		}
		g[i] = row
	}
	return g
}

func part1(data [][]int) int {
	lowPoints := findLowPoints(data)
	risk := 0
	for _, lp := range lowPoints {
		risk += data[lp.Y][lp.X] + 1
	}

	return risk
}

func part2(data [][]int) int {
	lowPoints := findLowPoints(data)
	checked := make(map[grid.Coordinate]struct{})
	basins := make(map[int]int)
	g := grid.Grid{
		X:      0,
		Y:      0,
		Width:  len(data[0]),
		Height: len(data),
	}
	for i, lp := range lowPoints {
		basins[i] = mapBasin(g, data, checked, lp)
	}

	// sort size of each basin, largest to smallest & multiply the top 3
	sizes := make([]int, 0, len(basins))
	for _, v := range basins {
		sizes = append(sizes, v)
		//fmt.Printf("%d - %d\n", b, v)
	}
	sort.Ints(sizes)
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	//fmt.Printf("%v\n", sizes)

	return sizes[0] * sizes[1] * sizes[2]
}

func mapBasin(g grid.Grid, data [][]int, checked map[grid.Coordinate]struct{}, c grid.Coordinate) int {
	count := 1
	checked[c] = struct{}{}
	for _, n := range c.Neighbors(false) {
		if _, ok := checked[n]; !ok && g.Valid(n) {
			cVal := data[c.Y][c.X]
			nVal := data[n.Y][n.X]
			if nVal != 9 && nVal >= cVal {
				count += mapBasin(g, data, checked, n)
			}
		}
	}
	return count
}

func findLowPoints(data [][]int) []grid.Coordinate {
	lowPoints := make([]grid.Coordinate, 0)
	g := grid.Grid{
		X:      0,
		Y:      0,
		Width:  len(data[0]),
		Height: len(data),
	}
	for y, row := range data {
		for x, v := range row {
			c := grid.Coordinate{X: x, Y: y}
			lowest := true
			for _, n := range c.Neighbors(false) {
				if g.Valid(n) {
					if v >= data[n.Y][n.X] {
						lowest = false
						break
					}
				}
			}
			if lowest {
				lowPoints = append(lowPoints, grid.Coordinate{
					X: x,
					Y: y,
				})
			}
		}
	}
	return lowPoints
}
