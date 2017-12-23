package days

import (
	"bufio"
	"fmt"

	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/input"
)

// Day13 is the 13th day in Advent of Code.
func Day13() *day.Day {
	readLayers := func() map[int]int {
		f := input.Open(13)
		scanner := bufio.NewScanner(f)
		layers := make(map[int]int)
		for scanner.Scan() {
			line := scanner.Text()
			var d, r int
			fmt.Sscanf(line, "%d: %d", &d, &r)
			layers[d] = r
		}
		return layers
	}

	part1 := func() interface{} {
		layers := readLayers()
		var penalty int
		for d, r := range layers {
			if d%((r-1)*2) == 0 {
				penalty += d * r
			}
		}

		return penalty
	}

	part2 := func() interface{} {
		layers := readLayers()

		pass := func(i int) bool {
			for d, r := range layers {
				if (d+i)%((r-1)*2) == 0 {
					return false
				}
			}
			return true
		}

		i := 0
		for {
			if pass(i) {
				return i
			}
			i++
		}
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.New(13, "Packet Scanners", solve)
}
