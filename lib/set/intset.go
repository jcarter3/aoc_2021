package set

import "sync"

type IntSet struct {
	l    sync.RWMutex
	data map[int]struct{}
}

func NewIntSet(els ...int) *IntSet {
	m := make(map[int]struct{})
	s := &IntSet{
		data: m,
	}
	s.Add(els...)
	return s
}

func (s *IntSet) Add(els ...int) {
	if len(els) == 0 {
		return
	}
	s.l.Lock()
	defer s.l.Unlock()
	for _, el := range els {
		s.data[el] = struct{}{}
	}
}

func (s *IntSet) Remove(els ...int) {
	if len(els) == 0 {
		return
	}
	s.l.Lock()
	defer s.l.Unlock()
	for _, el := range els {
		delete(s.data, el)
	}
}

func (s *IntSet) Has(e int) bool {
	_, ok := s.data[e]
	return ok
}

func (s *IntSet) List() []int {
	s.l.RLock()
	defer s.l.RUnlock()
	list := make([]int, 0, len(s.data))
	for el := range s.data {
		list = append(list, el)
	}
	return list
}

func (s *IntSet) Size() int {
	s.l.RLock()
	defer s.l.RUnlock()
	l := len(s.data)
	return l
}

func (s *IntSet) Separate(s2 *IntSet) {
	s.Remove(s2.List()...)
}

func (s *IntSet) Pop() int {
	s.l.RLock()

	for el := range s.data {
		s.l.RUnlock()
		s.l.Lock()
		delete(s.data, el)
		s.l.Unlock()
		return el
	}

	s.l.RUnlock()
	return 0
}
