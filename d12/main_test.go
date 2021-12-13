package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	data := Parse(lines)
	ans := part1(data)
	if ans != 10 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part1(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	data := Parse(lines)
	ans := part1(data)
	fmt.Printf("Answer: %d\n", ans)
	if ans != 4241 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	data := Parse(lines)
	ans := part2(data)
	if ans != 36 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part2(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	data := Parse(lines)
	ans := part2(data)
	fmt.Printf("intAnswer: %d\n", ans)
	if ans != 122_134 {
		t.Fail()
	}
}
