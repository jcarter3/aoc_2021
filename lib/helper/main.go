package helper

import (
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func BinaryStringToInt(s string) int {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func IntCompare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}
