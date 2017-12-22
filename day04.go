package main

import (
	"encoding/csv"
	"io"
	"os"
	"sort"
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

func day04() *day {

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
			check(err)

			valid := validate(record)

			if valid {
				count++
			}
		}

		return count
	}

	part1 := func() interface{} {
		f := openInput(4)
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
		f := openInput(4)
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

	return &day{4, "High-Entropy Passphrases", solve}
}
