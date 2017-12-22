package main

type stack struct {
	values []int
}

func (s *stack) pop() int {
	a := s.values
	x, a := a[0], a[1:]
	s.values = a
	return x
}

func (s *stack) push(xs ...int) {
	a := s.values
	a = append(a, xs...)
	s.values = a
}

func (s *stack) empty() bool {
	return len(s.values) == 0
}
