package main

import (
	"aoc_2021/lib/grid"
	"fmt"
	"log"
)

type Segment struct {
	x1, x2, y1, y2 int
}

func (s *Segment) IsStraight() bool {
	return s.x1 == s.x2 || s.y1 == s.y2
}

func (s *Segment) Coordinates() (grid.Coordinate, grid.Coordinate) {
	return grid.Coordinate{
			X: s.x1,
			Y: s.y1,
		}, grid.Coordinate{
			X: s.x2,
			Y: s.y2,
		}
}

func ParseData(lines []string) []*Segment {
	segments := make([]*Segment, 0)

	for _, l := range lines {
		s := &Segment{}
		_, err := fmt.Sscanf(l, "%d,%d -> %d,%d", &s.x1, &s.y1, &s.x2, &s.y2)
		if err != nil {
			log.Fatalf("failed to scan row (%s): %v", l, err)
		}
		segments = append(segments, s)
	}
	return segments
}

func part1(segments []*Segment) int {
	m := make(map[grid.Coordinate]int)

	for _, s := range segments {
		if !s.IsStraight() {
			continue
		}
		c1, c2 := s.Coordinates()
		path := c1.PathTo(c2)
		for _, c := range path {
			m[c]++
		}
	}
	count := 0
	for _, c := range m {
		if c > 1 {
			count++
		}
	}
	return count
}

func part2(segments []*Segment) int {
	m := make(map[grid.Coordinate]int)

	for _, s := range segments {
		c1, c2 := s.Coordinates()
		path := c1.PathTo(c2)
		for _, c := range path {
			m[c]++
		}
	}
	count := 0
	for _, c := range m {
		if c > 1 {
			count++
		}
	}
	return count
}
