package set

import (
	"fmt"
)

type Set[E comparable] struct {
	elements map[E]bool
}

func (s *Set[E]) Add(e E) {
	s.elements[e] = true
}

func (s *Set[E]) Remove(e E) {
	delete(s.elements, e)
}

func (s *Set[E]) Contains(e E) bool {
	_, ok := s.elements[e]

	return ok
}

func (s *Set[E]) Len() int {
	return len(s.elements)
}

func New[E comparable](elements ...E) Set[E] {
	set := Set[E]{
		elements: map[E]bool{},
	}
	for _, e := range elements {
		set.Add(e)
	}

	return set
}

func (s *Set[E]) String() string {
	var elements string
	size := len(s.elements)
	var cur int
	for k := range s.elements {
		if cur == size-1 {
			elements += fmt.Sprint(k)
		} else {
			elements += fmt.Sprint(k) + ", "
		}
	}

	return fmt.Sprintf("Set[%v]", elements)
}

func isNil[E comparable](arg E) bool {
	var t E
	return arg == t
}

func (s *Set[E]) ToSlice() []E {
	elements := make([]E, 0, len(s.elements))
	for e := range s.elements {
		if isNil(e) {
			continue
		}
		elements = append(elements, e)
	}

	return elements
}
