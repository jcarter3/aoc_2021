package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	dots, cuts := Parse(lines)
	ans := part1(dots, cuts)
	if ans != 17 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part1(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	dots, cuts := Parse(lines)
	ans := part1(dots, cuts)
	fmt.Printf("Answer: %d\n", ans)
	if ans != 755 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	dots, cuts := Parse(lines)
	ans := part2(dots, cuts)
	fmt.Printf("Ans:\n%s\n", ans) // O
}

func Test_Part2(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	dots, cuts := Parse(lines)
	ans := part2(dots, cuts)
	fmt.Printf("Answer:\n%s\n", ans) // BLKJRBAG
}
