package common

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Check panics if the given error is not nil.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// ParseInt parses an integer from the given string, ignoring whitespace and newlines.
// If the string does not contain an integer, it panics.
func ParseInt(str string) int {
	str = strings.Trim(str, "\r\t\n ")
	i, err := strconv.Atoi(str)
	Check(err)
	return i
}

// ReadInput reads the contents of the input file for the given day as a slice of bytes.
func ReadInput(day int) []byte {
	path := fmt.Sprintf("./input/day%d/input.txt", day)
	dat, err := ioutil.ReadFile(path)
	Check(err)
	return dat
}

// OpenInput opens the input file for the given day.
func OpenInput(day int) *os.File {
	path := fmt.Sprintf("./input/day%d/input.txt", day)
	f, err := os.Open(path)
	Check(err)
	return f
}

// ReadInputLine reads the first line of the input file for the given day, as a string.
func ReadInputLine(day int) string {
	file := OpenInput(day)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		return line
	}

	panic("Uh-oh!")
}
