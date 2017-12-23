package assert

// IsNil panics if the given value is not nil.
func IsNil(v interface{}) {
	if v != nil {
		panic(v)
	}
}

// IsTrue panics if the given value is not true.
func IsTrue(b bool) {
	if !b {
		panic(b)
	}
}
