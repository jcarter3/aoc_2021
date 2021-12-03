package main

import (
	"aoc_2021/lib/helper"
)

func part1(report []string) int {

	gamma := ""
	epsilon := ""
	for i := 0; i < len(report[0]); i++ {
		zeroes, ones := countsForPosition(report, i)
		if zeroes > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	g := helper.BinaryStringToInt(gamma)
	e := helper.BinaryStringToInt(epsilon)
	//fmt.Printf("Gamma: %s (%d)\nEpsilon: %s (%d)\n", gamma, g, epsilon, e)
	return g * e
}

func part2(report []string) int {
	oxygen := report
	co2 := report
	for i := 0; i < len(report[0]); i++ {
		zeroes, ones := countsForPosition(oxygen, i)
		if zeroes > ones {
			oxygen = filterReport(oxygen, i, '0')
		} else {
			oxygen = filterReport(oxygen, i, '1')
		}
		if len(oxygen) == 1 {
			break
		}
	}
	for i := 0; i < len(report[0]); i++ {
		zeroes, ones := countsForPosition(co2, i)
		if zeroes <= ones {
			co2 = filterReport(co2, i, '0')
		} else {
			co2 = filterReport(co2, i, '1')
		}
		if len(co2) == 1 {
			break
		}
	}
	o := helper.BinaryStringToInt(oxygen[0])
	c := helper.BinaryStringToInt(co2[0])
	//fmt.Printf("Oxygen: %s (%d)\nCO2: %s (%d)\n", oxygen[0], o, co2[0], c)
	return o * c
}

func countsForPosition(report []string, position int) (int, int) {
	zeroes, ones := 0, 0
	for _, r := range report {
		if r[position] == '0' {
			zeroes++
		} else {
			ones++
		}
	}
	return zeroes, ones
}

func filterReport(report []string, position int, val uint8) []string {
	filtered := make([]string, 0)
	for _, r := range report {
		if r[position] == val {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
