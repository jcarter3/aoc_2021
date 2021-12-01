package fileutil

import (
	"gitlab.com/jscarter3/aoc_go/common/helper"
	"io/ioutil"
	"strings"
)

func ReadLines(filename string) []string {
	return ReadSplit(filename, "\n")
}

func ReadSplit(filename, sep string) []string {
	return strings.Split(ReadAll(filename), sep)
}

func ReadSplitAsInt(filename, sep string) []int {
	raw := strings.Split(ReadAll(filename), sep)
	ret := make([]int, 0, len(raw))
	for _, s := range raw {
		ret = append(ret, helper.Atoi(s))
	}
	return ret
}

func ReadAll(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(b)
}
