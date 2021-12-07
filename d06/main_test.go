package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	data := fileutil.ReadSplitAsInt("sample.txt", ",")
	fishState := ParseData(data)
	ans := part1(fishState, 80)
	if ans != 5934 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part1(t *testing.T) {
	data := fileutil.ReadSplitAsInt("input.txt", ",")
	fishState := ParseData(data)
	ans := part1(fishState, 80)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 362740 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	data := fileutil.ReadSplitAsInt("sample.txt", ",")
	fishState := ParseData(data)
	ans := part1(fishState, 256) // same as part 1, just more days
	if ans != 26_984_457_539 {
		t.Fatalf("got %d", ans)
	}
}

func Test_Part2(t *testing.T) {
	data := fileutil.ReadSplitAsInt("input.txt", ",")
	fishState := ParseData(data)
	ans := part1(fishState, 256) // same as part 1, just more days
	fmt.Printf("Part 2 - Answer: %d\n", ans)
	if ans != 1_644_874_076_764 {
		t.Fail()
	}
}
