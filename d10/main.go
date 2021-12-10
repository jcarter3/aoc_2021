package main

import (
	"sort"
	"strings"
)

func part1(lines []string) int {
	score := 0
	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	for _, l := range lines {
		//fmt.Printf("processing line '%s'\n", l)
		incomplete, corrupted, next := ParseChunk(l)
		if incomplete {
			continue
		} else if corrupted {
			score += points[next]
		}
	}
	return score
}

func part2(lines []string) int {
	scores := make([]int, 0)
	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	for _, l := range lines {
		score := 0
		//fmt.Printf("processing line '%s'\n", l)
		_, corrupted, next := ParseChunk(l)
		if corrupted {
			continue
		}
		for _, s := range strings.Split(next, "") {
			score *= 5
			score += points[s]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)

	mid := (len(scores) / 2)
	return scores[mid]
}

func ParseChunk(chunk string) (bool, bool, string) {
	openClosers := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	inChunk := make([]string, 0)
	level := 0
	for _, s := range strings.Split(chunk, "") {
		//indent := strings.Repeat("  ", level)
		if _, ok := openClosers[s]; ok {
			inChunk = append([]string{s}, inChunk...)
			//fmt.Printf("%s%s\n", indent, s)
			level++
		} else if s == openClosers[inChunk[0]] { // is valid closer
			inChunk = inChunk[1:]
			//fmt.Printf("%s%s\n", indent, s)
			level--
		} else {
			return false, true, s
		}
	}
	next := ""
	for _, n := range inChunk {
		next += openClosers[n]
	}
	return true, false, next
}
