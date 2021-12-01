package set

import "sync"

type Set[T comparable] struct {
	l    sync.RWMutex
	data map[T]struct{}
}

func New[T comparable](els ...T) *Set[T] {
	m := make(map[T]struct{})
	s := &Set[T]{
		data: m,
	}
	s.Add(els...)
	return s
}

func (s *Set[T]) Add(els ...T) {
	if len(els) == 0 {
		return
	}
	s.l.Lock()
	defer s.l.Unlock()
	for _, el := range els {
		s.data[el] = struct{}{}
	}
}

func (s *Set[T]) Remove(els ...T) {
	if len(els) == 0 {
		return
	}
	s.l.Lock()
	defer s.l.Unlock()
	for _, el := range els {
		delete(s.data, el)
	}
}

func (s *Set[T]) Has(e T) bool {
	_, ok := s.data[e]
	return ok
}

func (s *Set[T]) List() []T {
	s.l.RLock()
	defer s.l.RUnlock()
	list := make([]T, 0, len(s.data))
	for el := range s.data {
		list = append(list, el)
	}
	return list
}

func (s *Set[T]) Size() int {
	s.l.RLock()
	defer s.l.RUnlock()
	l := len(s.data)
	return l
}

func (s *Set[T]) Separate(s2 *Set[T]) {
	s.Remove(s2.List()...)
}

func (s *Set[T]) Pop() interface{} {
	s.l.RLock()

	for el := range s.data {
		s.l.RUnlock()
		s.l.Lock()
		delete(s.data, el)
		s.l.Unlock()
		return el
	}

	s.l.RUnlock()
	return nil
}
