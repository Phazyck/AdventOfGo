package main

import "testing"

func assert(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("expected '%v' got '%v'", expected, actual)
	}
}

// TestExample1 tests that solveInput works for the 1st example
func TestExample1(t *testing.T) {
	result, _ := solveInput("ne,ne,ne")
	assert(t, result, 3)
}

// TestExample2 tests that solveInput works for the 2nd example
func TestExample2(t *testing.T) {
	result, _ := solveInput("ne,ne,sw,sw")
	assert(t, result, 0)
}

// TestExample3 tests that solveInput works for the 3rd example
func TestExample3(t *testing.T) {
	result, _ := solveInput("ne,ne,s,s")
	assert(t, result, 2)
}

// TestExample4 tests that solveInput works for the 4th example
func TestExample4(t *testing.T) {
	result, _ := solveInput("se,sw,se,sw,sw")
	assert(t, result, 3)
}
