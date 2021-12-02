package d02

import (
	"aoc_2021/lib/fileutil"
	"fmt"
	"testing"
)

func Test_Part1(t *testing.T) {
	instructions := fileutil.ReadLines("input.txt")
	ans := part1(instructions)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 1604850 {
		t.Fail()
	}
}

func Test_Part2(t *testing.T) {
	instructions := fileutil.ReadLines("input.txt")
	ans := part2(instructions)
	fmt.Printf("Part 2 - Answer: %d\n", ans)
	if ans != 1685186100 {
		t.Fail()
	}
}
