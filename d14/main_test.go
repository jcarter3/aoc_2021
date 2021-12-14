package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	template, rules := Parse(lines)
	ans := part1(template, rules, 10)
	if ans != 1588 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part1(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	template, rules := Parse(lines)
	ans := part1(template, rules, 10)
	fmt.Printf("Answer: %d\n", ans)
	if ans != 3831 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	lines := fileutil.ReadLines("sample.txt")
	template, rules := Parse(lines)
	ans := part1(template, rules, 40)
	if ans != 2_188_189_693_529 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part2(t *testing.T) {
	lines := fileutil.ReadLines("input.txt")
	template, rules := Parse(lines)
	ans := part1(template, rules, 40)
	fmt.Printf("Answer: %d\n", ans)
	if ans != 5_725_739_914_282 {
		t.Fail()
	}
}
