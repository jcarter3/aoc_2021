package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	grid := ParseGrid(lines)
	ans := part1(grid)
	if ans != 15 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part1(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	grid := ParseGrid(lines)
	ans := part1(grid)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 502 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	grid := ParseGrid(lines)
	ans := part2(grid)
	if ans != 1134 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part2(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	grid := ParseGrid(lines)
	ans := part2(grid)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 1_330_560 {
		t.Fail()
	}
}
