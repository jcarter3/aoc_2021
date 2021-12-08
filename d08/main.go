package main

import (
	"aoc_2021/lib/helper"
	"aoc_2021/lib/set"
	"log"
	"sort"
	"strings"
)

type Entry struct {
	NumToPattern   map[int]string
	PatternToNum   map[string]string
	SignalPatterns [10]string
	Output         [4]string
}

func (e *Entry) Decode() {
	e.NumToPattern = make(map[int]string)
	e.PatternToNum = make(map[string]string)
	easyMapped := 0
	unmapped := set.New[string](e.SignalPatterns[:]...)
	for unmapped.Size() > 0 {
		for _, s := range unmapped.List() {
			if len(s) == 2 {
				e.NumToPattern[1] = s
				e.PatternToNum[s] = "1"
				//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
				easyMapped++
				unmapped.Remove(s)
			} else if len(s) == 4 {
				e.NumToPattern[4] = s
				e.PatternToNum[s] = "4"
				//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
				easyMapped++
				unmapped.Remove(s)
			} else if len(s) == 3 {
				e.NumToPattern[7] = s
				e.PatternToNum[s] = "7"
				//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
				easyMapped++
				unmapped.Remove(s)
			} else if len(s) == 7 {
				e.NumToPattern[8] = s
				e.PatternToNum[s] = "8"
				//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
				easyMapped++
				unmapped.Remove(s)
			}
			if easyMapped < 4 {
				continue // don't continue until 1, 4, 7 and 8 are all mapped
			}

			//fmt.Printf("attempting to map: %s\n", s)
			if len(s) == 6 { // either 0, 6 or 9
				if contains(s, e.NumToPattern[4]) { // 9 contains all of 4
					e.NumToPattern[9] = s
					e.PatternToNum[s] = "9"
					//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
					unmapped.Remove(s)
				} else {
					if contains(s, e.NumToPattern[1]) {
						e.NumToPattern[0] = s
						e.PatternToNum[s] = "0"
						//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
						unmapped.Remove(s)
					} else {
						e.NumToPattern[6] = s
						e.PatternToNum[s] = "6"
						//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
						unmapped.Remove(s)
					}
				}
			} else if len(s) == 5 { // length 5, either 2, 3, or 5
				if v9, is9Mapped := e.NumToPattern[9]; is9Mapped {
					if contains(v9, s) { // contained by 9, so is either 3 or 5
						if contains(s, e.NumToPattern[1]) { // if contains 1, is 3. else is 5
							e.NumToPattern[3] = s
							e.PatternToNum[s] = "3"
							//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
							unmapped.Remove(s)
						} else {
							e.NumToPattern[5] = s
							e.PatternToNum[s] = "5"
							//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
							unmapped.Remove(s)
						}
					} else {
						e.NumToPattern[2] = s
						e.PatternToNum[s] = "2"
						//fmt.Printf("mapping %s to %s\n", s, e.PatternToNum[s])
						unmapped.Remove(s)
					}
				}
			}
		}
	}
}

func ParseEntries(lines []string) []Entry {
	entries := make([]Entry, 0, len(lines))
	for _, l := range lines {
		entry := Entry{}
		patterns, outputValue, _ := strings.Cut(l, " | ")
		for i, s := range strings.Split(patterns, " ") {
			entry.SignalPatterns[i] = sortString(s)
		}
		for i, s := range strings.Split(outputValue, " ") {
			entry.Output[i] = sortString(s)
		}
		entry.Decode()
		entries = append(entries, entry)
	}
	return entries
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func contains(a, b string) bool {
	s1 := set.New[string](strings.Split(a, "")...)
	return s1.Has(strings.Split(b, "")...)
}

func part1(entries []Entry) int {
	uniqueNums := set.NewIntSet(2, 4, 3, 7) // 1, 4, 7, 8
	count := 0
	for _, entry := range entries {
		for _, v := range entry.Output {
			if uniqueNums.Has(len(v)) {
				count++
			}
		}
	}

	return count
}

func part2(entries []Entry) int {
	sum := 0
	for _, entry := range entries {
		s := ""
		for _, v := range entry.Output {
			if num, ok := entry.PatternToNum[v]; ok {
				s += num
			} else {
				log.Fatalf("this shouldn't happen")
			}
		}
		//fmt.Printf("%s - %s\n", strings.Join(entry.Output[:], " "), s)
		sum += helper.Atoi(s)
	}
	return sum
}
