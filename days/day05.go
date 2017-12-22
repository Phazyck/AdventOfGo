package days

import (
	"bufio"
	"os"

	"github.com/Phazyck/AdventOfGo/common"
	"github.com/Phazyck/AdventOfGo/day"
)

// Day05 is the 5th day in Advent of Code.
func Day05() *day.Day {

	traverse := func(list []int, exec func([]int, int) int) int {
		i := 0
		steps := 0

		for {
			if i < 0 || i >= len(list) {
				break
			}
			i = exec(list, i)
			steps++
		}

		return steps
	}

	readInts := func(f *os.File) []int {
		ints := make([]int, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()

			i := common.ParseInt(line)

			ints = append(ints, i)
		}

		return ints
	}

	part1 := func() interface{} {
		f := common.OpenInput(5)
		list := readInts(f)

		exec := func(list []int, i int) int {
			offset := list[i]
			list[i] = offset + 1
			return i + offset
		}

		return traverse(list, exec)
	}

	part2 := func() interface{} {
		f := common.OpenInput(5)
		list := readInts(f)

		exec := func(list []int, i int) int {
			offset := list[i]
			if offset >= 3 {
				list[i] = offset - 1
			} else {
				list[i] = offset + 1
			}

			return i + offset
		}

		return traverse(list, exec)
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.NewDay(5, "A Maze of Twisty Trampolines, All Alike", solve)
}
