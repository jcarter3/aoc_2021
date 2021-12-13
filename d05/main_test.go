package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	data := fileutil.ReadLines("sample.txt")
	segments := ParseData(data)
	ans := part1(segments)
	if ans != 5 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part1(t *testing.T) {
	data := fileutil.ReadLines("input.txt")
	segments := ParseData(data)
	ans := part1(segments)
	fmt.Printf("Answer: %d\n", ans)
	if ans != 7644 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	data := fileutil.ReadLines("sample.txt")
	segments := ParseData(data)
	ans := part2(segments)
	if ans != 12 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part2(t *testing.T) {
	data := fileutil.ReadLines("input.txt")
	segments := ParseData(data)
	ans := part2(segments)
	fmt.Printf("intAnswer: %d\n", ans)
	if ans != 18627 {
		t.Fail()
	}
}
