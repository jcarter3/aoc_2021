package main

func ParseData(fish []int) map[int]int {
	fishState := make(map[int]int)
	for _, f := range fish {
		fishState[f]++
	}

	return fishState
}

func part1(fs map[int]int, days int) int {
	for i := 0; i < days; i++ {
		newData := make(map[int]int)
		for k, v := range fs {
			if k == 0 {
				newData[6] += v
				newData[8] += v
			} else {
				newData[k-1] += v
			}
		}
		fs = newData
	}

	totalFish := 0
	for _, v := range fs {
		totalFish += v
	}
	return totalFish
}
