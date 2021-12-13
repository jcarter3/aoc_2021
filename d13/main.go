package main

import (
	"aoc_2021/lib/grid"
	"aoc_2021/lib/helper"
	"bytes"
	"github.com/icholy/draw"
	"log"
	"math"
	"strings"
)

func part1(dots map[grid.Coordinate]struct{}, cuts []string) int {
	dots = Fold(dots, cuts[0])

	return len(dots)
}

func part2(dots map[grid.Coordinate]struct{}, cuts []string) string {
	for _, cut := range cuts {
		dots = Fold(dots, cut)
	}
	maxX, maxY := math.MinInt, math.MinInt
	for d := range dots {
		if d.X > maxX {
			maxX = d.X
		}
		if d.Y > maxY {
			maxY = d.Y
		}
	}
	cv := draw.NewCanvas(maxX+1, maxY+1)
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			c := grid.Coordinate{X: x, Y: y}
			_, ok := dots[c]
			var char byte = '.'
			if ok {
				char = '#'
			}
			cv.Draw(draw.Point{X: float64(x), Y: float64(y)}, char)
		}
	}
	var buf bytes.Buffer
	if err := cv.WriteTo(&buf); err != nil {
		log.Fatalf("failed to write buffer: %v", err)
	}
	return buf.String()
}

func Fold(dots map[grid.Coordinate]struct{}, cut string) map[grid.Coordinate]struct{} {
	dir, placeStr, _ := strings.Cut(cut, "=")
	place := helper.Atoi(placeStr)

	folded := make(map[grid.Coordinate]struct{})
	for dot := range dots {
		if dir == "x" {
			if dot.X < place {
				folded[dot] = struct{}{}
			} else {
				newDot := grid.Coordinate{
					X: place - (dot.X - place),
					Y: dot.Y,
				}
				folded[newDot] = struct{}{}
			}
		} else {
			if dot.Y < place {
				folded[dot] = struct{}{}
			} else {
				newDot := grid.Coordinate{
					X: dot.X,
					Y: place - (dot.Y - place),
				}
				folded[newDot] = struct{}{}
			}
		}
	}
	return folded
}

func Parse(lines []string) (map[grid.Coordinate]struct{}, []string) {
	dots := make(map[grid.Coordinate]struct{})
	var cuts []string
	for _, line := range lines {
		if strings.HasPrefix(line, "fold along ") {
			cuts = append(cuts, strings.TrimPrefix(line, "fold along "))
		} else if line == "" {
			continue
		} else {
			x, y, _ := strings.Cut(line, ",")
			c := grid.Coordinate{X: helper.Atoi(x), Y: helper.Atoi(y)}
			dots[c] = struct{}{}
		}
	}
	return dots, cuts
}
