package days

import (
	"strings"

	"github.com/Phazyck/AdventOfGo/common"
	"github.com/Phazyck/AdventOfGo/day"
)

func solveInput(input string) (interface{}, interface{}) {
	steps := strings.Split(input, ",")

	l := 6
	taken := make([]int, l, l)

	toIndex := map[string]int{
		"n":  0,
		"ne": 1,
		"se": 2,
		"s":  3,
		"sw": 4,
		"nw": 5,
	}

	takeStep := func(step string) int {
		idx := toIndex[step]

		// negation
		rev := (idx + 3) % l
		if taken[rev] > 0 {
			taken[rev]--
			return -1
		}

		// shortcut a
		ir := (idx + 1) % l
		irr := (idx + 2) % l

		if taken[irr] > 0 {
			taken[irr]--
			taken[ir]++
			return 0
		}

		// shortcut b
		il := (idx + l - 1) % l
		ill := (idx + l - 2) % l

		if taken[ill] > 0 {
			taken[ill]--
			taken[il]++
			return 0
		}

		// regular walk
		taken[idx]++
		return 1
	}

	var count, max int
	for _, step := range steps {
		delta := takeStep(step)
		count += delta
		if count > max {
			max = count
		}
	}

	return count, max
}

// Day11 is the 11th day in Advent of Code.
func Day11() *day.Day {
	solve := func() (interface{}, interface{}) {
		input := common.ReadInputLine(11)
		part1, part2 := solveInput(input)
		return part1, part2
	}

	return day.NewDay(11, "Hex Ed", solve)
}
