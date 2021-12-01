package d01

func part1(depths []int) int {
	return countIncreases(depths)
}

func part2(depths []int) int {
	windows := make([]int, 0)
	for i := 0; i < len(depths)-2; i++ {
		sum := depths[i] + depths[i+1] + depths[i+2]
		windows = append(windows, sum)
	}
	return countIncreases(windows)
}

func countIncreases(nums []int) int {
	increases := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			increases++
		}
	}
	return increases
}
