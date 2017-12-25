package input

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/phazyck/adventofgo/assert"
)

func getPath(day int) string {
	_, f, _, ok := runtime.Caller(0)
	assert.IsTrue(ok)
	dir := path.Dir(f)
	path := path.Join(dir, fmt.Sprintf("day%02d", day), "input.txt")
	return path
}

// Read reads the contents of the input file for the given day as a slice of bytes.
func Read(day int) []byte {
	path := getPath(day)
	dat, err := ioutil.ReadFile(path)
	assert.IsNil(err)
	return dat
}

// Open opens the input file for the given day.
func Open(day int) *os.File {
	path := getPath(day)
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

// ReadLines reads the first n lines of the input file for the given day, as a slice of strings.
func ReadLines(day, n int) []string {
	file := Open(day)
	scanner := bufio.NewScanner(file)
	lines := make([]string, n)

	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
		if i == n {
			return lines
		}
	}

	panic("Uh-oh!")
}
