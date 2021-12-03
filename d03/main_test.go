package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	report := fileutil.ReadLines("sample.txt")
	ans := part1(report)
	if ans != 198 {
		t.Fail()
	}
}

func Test_Part1(t *testing.T) {
	report := fileutil.ReadLines("input.txt")
	ans := part1(report)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 3969000 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	report := fileutil.ReadLines("sample.txt")
	ans := part2(report)
	fmt.Printf("Part 2 - Answer: %d\n", ans)
	if ans != 230 {
		t.Fail()
	}
}

func Test_Part2(t *testing.T) {
	report := fileutil.ReadLines("input.txt")
	ans := part2(report)
	fmt.Printf("Part 2 - Answer: %d\n", ans)
	if ans != 4267809 {
		t.Fail()
	}
}
