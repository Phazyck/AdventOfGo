package test

import "testing"

// AssertEqual fails the given test is the actual value does not equal the expected value.
func AssertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("expected '%v' got '%v'", expected, actual)
	}
}
