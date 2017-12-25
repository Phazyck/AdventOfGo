package days

import (
	"encoding/csv"
	"io"
	"os"
	"sort"

	"github.com/phazyck/adventofgo/assert"
	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
)

type sortableString []rune

func (s sortableString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortableString) Len() int {
	return len(s)
}

// Day04 is the 4th day in Advent of Code.
func Day04() *day.Day {

	type validator func([]string) bool

	countValid := func(f *os.File, validate validator) int {
		r := csv.NewReader(f)
		r.Comma = ' '
		r.FieldsPerRecord = -1

		count := 0
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			assert.IsNil(err)

			valid := validate(record)

			if valid {
				count++
			}
		}

		return count
	}

	part1 := func() interface{} {
		f := input.Open(4)
		validate := func(words []string) bool {
			used := make(map[string]bool)

			for _, word := range words {

				if used[word] {
					return false
				}
				used[word] = true
			}

			return true
		}
		return countValid(f, validate)
	}

	part2 := func() interface{} {
		f := input.Open(4)
		validate := func(words []string) bool {
			used := make(map[string]bool)

			for _, word := range words {
				sortable := sortableString(word)
				sort.Sort(sortable)
				word = string(sortable)
				if used[word] {
					return false
				}
				used[word] = true
			}

			return true
		}
		return countValid(f, validate)
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.New(4, "High-Entropy Passphrases", solve)
}
