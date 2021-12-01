package helper

import "encoding/hex"

func KnotHash(input string) []byte {
	lengths := make([]int, 0, len(input)+5)
	for _, r := range input {
		lengths = append(lengths, int(r))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)
	list := initList()
	pos := 0
	skip := 0
	for i := 0; i < 64; i++ {
		pos, skip = process(list, lengths, pos, skip)
	}
	return densify(list)
}

func KnotHashString(input string) string {
	return hex.EncodeToString(KnotHash(input))
}

func initList() []int {
	ints := make([]int, 256)
	for i := range ints {
		ints[i] = i
	}
	return ints
}

func process(list, lengths []int, pos, skip int) (int, int) {
	for _, l := range lengths {
		reverse(list, pos, l)
		pos = (pos + l + skip) % len(list)
		skip++
	}
	return pos, skip
}

func reverse(list []int, pos, length int) {
	double := append(list, list...)
	for i, j := pos, pos+length-1; i < j; i, j = i+1, j-1 {
		double[i], double[j] = double[j], double[i]
	}

	copy(list[pos:], double[pos:pos+length])
	if overlap := (pos + length) - len(list); overlap > 0 {
		copy(list[:overlap], double[len(list):])
	}
}

func densify(list []int) []byte {
	dense := make([]byte, len(list)/16)
	for i := range dense {
		d := 0
		for _, l := range list[i*16:][:16] {
			d ^= l
		}
		dense[i] = byte(d)
	}
	return dense
}
