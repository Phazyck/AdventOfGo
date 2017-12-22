package input

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Phazyck/AdventOfGo/assert"
)

// Read reads the contents of the input file for the given day as a slice of bytes.
func Read(day int) []byte {
	path := fmt.Sprintf("./input/day%d/input.txt", day)
	dat, err := ioutil.ReadFile(path)
	assert.IsNil(err)
	return dat
}

// Open opens the input file for the given day.
func Open(day int) *os.File {
	path := fmt.Sprintf("./input/day%d/input.txt", day)
	f, err := os.Open(path)
	assert.IsNil(err)
	return f
}

// ReadLine reads the first line of the input file for the given day, as a string.
func ReadLine(day int) string {
	file := Open(day)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		return line
	}

	panic("Uh-oh!")
}
