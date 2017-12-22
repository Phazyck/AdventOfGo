package stack

// Stack is a stack of integers.
type Stack struct {
	values []int
}

// New returns a new stack with the given values.
func New(values ...int) *Stack {
	return &Stack{values}
}

// Pop takes an integer from the top of the stack.
func (s *Stack) Pop() int {
	a := s.values
	x, a := a[0], a[1:]
	s.values = a
	return x
}

// Push places the given integers on the top of the stack.
func (s *Stack) Push(xs ...int) {
	a := s.values
	a = append(a, xs...)
	s.values = a
}

// Empty returns true if the stack contains no integers, otherwise false.
func (s *Stack) Empty() bool {
	return len(s.values) == 0
}
