package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	entries := ParseEntries(lines)
	ans := part1(entries)
	if ans != 26 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part1(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	entries := ParseEntries(lines)
	ans := part1(entries)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 367 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	entries := ParseEntries(lines)
	ans := part2(entries)
	if ans != 61229 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part2(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	entries := ParseEntries(lines)
	ans := part2(entries)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 974_512 {
		t.Fail()
	}
}
