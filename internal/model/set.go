package model

type Set struct {
	elements map[int]struct{}
}

func NewSet() *Set {
	return &Set{elements: make(map[int]struct{})}
}

func (s *Set) Add(value int) {
	s.elements[value] = struct{}{}
}

func (s *Set) Remove(value int) {
	delete(s.elements, value)
}

func (s *Set) Contains(value int) bool {
	_, exists := s.elements[value]
	return exists
}

func (s *Set) Elements() []int {
	keys := make([]int, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}
