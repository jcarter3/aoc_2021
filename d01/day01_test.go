package d01

import (
	"fmt"
	"testing"

	"gitlab.com/jscarter3/aoc_go/common/fileutil"
)

func Test_Sample1_1(t *testing.T) {
	ans := part1([]int{1721, 979, 366, 299, 675, 1456})
	if ans != 514579 {
		t.Fatalf("part1 sample: expected: %d, actual: %d", 514579, ans)
	}
}

func Test_Part1(t *testing.T) {
	expenses := fileutil.ReadSplitAsInt("input.txt", "\n")
	ans := part1(expenses)
	fmt.Printf("Part 1 - Answer: %d\n", ans)
	if ans != 252724 {
		t.Fail()
	}
}

func Test_Sample2_1(t *testing.T) {
	ans := part2([]int{1721, 979, 366, 299, 675, 1456})
	if ans != 241861950 {
		t.Fatalf("part2 sample: expected: %d, actual: %d", 514579, ans)
	}
}

func Test_Part2(t *testing.T) {
	expenses := fileutil.ReadSplitAsInt("input.txt", "\n")
	ans := part2(expenses)
	fmt.Printf("Part 2 - Answer: %d\n", ans)
	if ans != 276912720 {
		t.Fail()
	}
}
