package assert

// IsNil panics if the given value is not nil.
func IsNil(v interface{}) {
	if v != nil {
		panic(v)
	}
}
