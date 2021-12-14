package main

import (
	"aoc_2021/lib/counter"
	"strings"
)

func part1(template string, rules map[string]string, steps int) int {
	pairs := make(map[string]int)
	for i := 1; i < len(template); i++ {
		pairs[template[i-1:i+1]]++
	}

	for step := 0; step < steps; step++ {
		newPairs := make(map[string]int)
		for pair, count := range pairs {
			r := rules[pair]
			k1 := pair[0:1] + r
			k2 := r + pair[1:2]
			newPairs[k1] += count
			newPairs[k2] += count
		}
		pairs = newPairs
	}
	rc := counter.NewRuneCounter()
	for pair, count := range pairs {
		rc.AddRuneCount(rune(pair[0]), count)
	}
	rc.AddRune(rune(template[len(template)-1]))
	runes := rc.Sort()
	return rc.Counts[runes[0]] - rc.Counts[runes[len(runes)-1]]
}

func Parse(lines []string) (string, map[string]string) {
	insertionRules := make(map[string]string)
	for _, line := range lines[2:] {
		pair, element, _ := strings.Cut(line, " -> ")
		insertionRules[pair] = element
	}
	return lines[0], insertionRules
}
