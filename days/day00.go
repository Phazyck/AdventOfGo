package days

import "github.com/phazyck/adventofgo/day"

// Day00 is the 00th day in Advent of Code.
// (Actually, it's a template for implementing days.)
func Day00() *day.Day {

	part1 := func() interface{} {
		return -1
	}

	part2 := func() interface{} {
		return -1
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.New(00, "Title", solve)
}
