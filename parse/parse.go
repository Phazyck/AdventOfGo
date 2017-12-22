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
