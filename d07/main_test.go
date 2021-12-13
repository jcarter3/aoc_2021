package main

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Sample1(t *testing.T) {
	data := fileutil.ReadSplitAsInt("sample.txt", ",")
	position, fuel := part1(data)
	if fuel != 37 {
		t.Fatalf("got %d & %d", position, fuel)
	}
}

func Test_Part1(t *testing.T) {
	data := fileutil.ReadSplitAsInt("input.txt", ",")
	position, fuel := part1(data)
	fmt.Printf("Position: %d, Fuel: %d\n", position, fuel)
	if fuel != 345_035 {
		t.Fail()
	}
}

func Test_Sample2(t *testing.T) {
	data := fileutil.ReadSplitAsInt("sample.txt", ",")
	position, fuel := part2(data)
	if fuel != 168 {
		t.Fatalf("got %d & %d", position, fuel)
	}
}

func Test_Part2(t *testing.T) {
	data := fileutil.ReadSplitAsInt("input.txt", ",")
	position, fuel := part2(data)
	fmt.Printf("Position: %d, Fuel: %d\n", position, fuel)
	if fuel != 97_038_163 {
		t.Fail()
	}
}
