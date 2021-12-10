package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	ans := part1(lines)
	if ans != 26397 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part1(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	ans := part1(lines)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 366027 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	ans := part2(lines)
	if ans != 288957 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part2(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	ans := part2(lines)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 1_118_645_287 {
		t.Fail()
	}
}
