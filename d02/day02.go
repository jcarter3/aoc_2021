package d02

import (
	"aoc_2021/lib/helper"
	"strings"
)

func part1(instructions []string) int {
	horizontal, depth := 0, 0
	for _, instruction := range instructions {
		direction, changeStr, _ := strings.Cut(instruction, " ")
		change := helper.Atoi(changeStr)
		switch direction {
		case "forward":
			horizontal += change
		case "down":
			depth += change
		case "up":
			depth -= change
		}
	}
	return horizontal * depth
}

func part2(instructions []string) int {
	horizontal, depth, aim := 0, 0, 0
	for _, instruction := range instructions {
		direction, changeStr, _ := strings.Cut(instruction, " ")
		change := helper.Atoi(changeStr)
		switch direction {
		case "forward":
			horizontal += change
			depth += aim * change
		case "down":
			aim += change
		case "up":
			aim -= change
		}
	}
	return horizontal * depth
}
