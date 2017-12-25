package stack

// Stack is a stack of integers.
type Stack []int

// New returns a new stack with the given values.
func New(vals ...int) *Stack {
	stk := Stack(vals)
	return &stk
}

// Pop takes an integer from the top of the stack.
func (stk *Stack) Pop() int {
	arr := *stk
	val, arr := arr[0], arr[1:]
	*stk = arr
	return val
}

// Push places the given integers on the top of the stack.
func (stk *Stack) Push(vals ...int) {
	arr := *stk
	arr = append(arr, vals...)
	*stk = arr
}

// Empty returns true if the stack contains no integers, otherwise false.
func (stk *Stack) Empty() bool {
	arr := *stk
	return len(arr) == 0
}
