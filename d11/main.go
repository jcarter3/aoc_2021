package main

import (
	"aoc_2021/lib/grid"
	"aoc_2021/lib/helper"
	"strings"
)

func part1(octopi *Octopi, steps int) int {
	for i := 0; i < steps; i++ {
		octopi.Step()
	}
	return octopi.flashes
}

func part2(octopi *Octopi) int {
	lastFlashes := 0
	steps := 0
	for {
		steps++
		octopi.Step()
		if octopi.flashes-lastFlashes == 100 {
			break
		} else {
			lastFlashes = octopi.flashes
		}
	}
	return steps
}

type Octopi struct {
	data    [10][10]int
	flashes int
}

func (o *Octopi) Step() {
	// Increment all by 1
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			o.data[y][x]++
		}
	}
	for {
		newFlashes := 0
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				v := o.data[y][x]
				c := grid.Coordinate{X: x, Y: y}
				if v > 9 {
					newFlashes++
					o.data[y][x] = 0
					for _, n := range c.Neighbors(true) {
						if o.ValidCoord(n.X, n.Y) && o.data[n.Y][n.X] > 0 {
							o.data[n.Y][n.X]++
						}
					}
				}
			}
		}
		if newFlashes == 0 {
			break
		}
		o.flashes += newFlashes
	}
}

func (o *Octopi) ValidCoord(x, y int) bool {
	return x >= 0 && y >= 0 && x < 10 && y < 10
}

func Parse(lines []string) *Octopi {
	var data [10][10]int
	for y, line := range lines {
		for x, s := range strings.Split(line, "") {
			data[y][x] = helper.Atoi(s)
		}
	}
	return &Octopi{
		data: data,
	}
}
