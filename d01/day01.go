package d01

func part1(expenses []int) int {
	for i := 0; i < len(expenses)-1; i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				return expenses[i] * expenses[j]
			}
		}
	}
	panic("didn't find answer")
}

func part2(expenses []int) int {
	for i := 0; i < len(expenses)-2; i++ {
		for j := i + 1; j < len(expenses)-1; j++ {
			for k := j + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return expenses[i] * expenses[j] * expenses[k]
				}
			}
		}
	}
	panic("didn't find answer")
}
