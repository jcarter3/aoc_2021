package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Part1(t *testing.T) {
	depths := fileutil.ReadSplitAsInt("input.txt", "\n")
	ans := part1(depths)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 1316 {
		t.Fail()
	}
}

func Test_Part2(t *testing.T) {
	depths := fileutil.ReadSplitAsInt("input.txt", "\n")
	ans := part2(depths)
	fmt.Printf("Part 2 - Answer: %d\n", ans)
	if ans != 1344 {
		t.Fail()
	}
}
