package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	data := fileutil.ReadLines("sample.txt")
	drawn, boards := ParseData(data)
	ans := part1(drawn, boards)
	if ans != 4512 {
		t.Fail()
	}
}

func Test_Part1(t *testing.T) {
	data := fileutil.ReadLines("input.txt")
	drawn, boards := ParseData(data)
	ans := part1(drawn, boards)
	fmt.Printf("Answer: %d\n", ans)
	if ans != 29440 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	data := fileutil.ReadLines("sample.txt")
	drawn, boards := ParseData(data)
	ans := part2(drawn, boards)
	if ans != 1924 {
		t.Fail()
	}
}

func Test_Part2(t *testing.T) {
	data := fileutil.ReadLines("input.txt")
	drawn, boards := ParseData(data)
	ans := part2(drawn, boards)
	fmt.Printf("intAnswer: %d\n", ans)
	if ans != 13884 {
		t.Fail()
	}
}
