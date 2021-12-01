package counter

import "sort"

type RuneCounter struct {
	Counts map[rune]int
}

func NewRuneCounter() *RuneCounter {
	return &RuneCounter{
		Counts: make(map[rune]int),
	}
}

func (rc *RuneCounter) AddString(s string) {
	for _, r := range s {
		rc.Counts[r]++
	}
}

func (rc *RuneCounter) AddRune(r rune) {
	rc.Counts[r]++
}

func (rc *RuneCounter) Sort() []rune {
	type pair struct {
		r rune
		n int
	}
	pairs := make([]pair, 0, len(rc.Counts))
	for r, n := range rc.Counts {
		pairs = append(pairs, pair{r, n})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].n == pairs[j].n {
			return pairs[i].r < pairs[j].r
		}
		return pairs[i].n > pairs[j].n
	})
	ret := make([]rune, 0, len(pairs))
	for _, r := range pairs {
		ret = append(ret, r.r)
	}
	return ret
}
