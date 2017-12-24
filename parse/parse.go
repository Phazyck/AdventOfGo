package parse

import (
	"strconv"
	"strings"

	"github.com/Phazyck/AdventOfGo/assert"
)

// AssertInt parses an integer from the given string, ignoring whitespace and newlines.
// If the string does not contain an integer, it panics.
func AssertInt(str string) int {
	str = strings.Trim(str, "\r\t\n ")
	i, err := strconv.Atoi(str)
	assert.IsNil(err)
	return i
}

// AssertInts parses a slice of integers a string of sub-strings, separated by a given separator.
// If any sub-string does not contain an integer, it panics.
func AssertInts(str, separator string) []int {
	strs := strings.Split(str, separator)
	count := len(strs)
	ints := make([]int, count, count)
	for i, s := range strs {
		ints[i] = AssertInt(s)
	}
	return ints
}

// AssertChar parses a single char from the given string.
// If the string contains anything but a single char, it will panic.
func AssertChar(str string) byte {
	if len(str) != 1 {
		panic(str)
	}

	return str[0]
}

// AssertByte parses a byte from the given string, ignoring whitespace and newlines.
// If the string does not contain an byte, it panics.
func AssertByte(str string) byte {
	str = strings.Trim(str, "\r\t\n ")
	i, err := strconv.Atoi(str)
	assert.IsNil(err)
	b := byte(i)
	if int(b) != i {
		panic("integer does not fit in a byte")
	}
	return b
}
