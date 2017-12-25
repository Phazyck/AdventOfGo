package days

import (
	"bufio"
	"fmt"

	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
)

// Day13 is the 13th day in Advent of Code.
func Day13() *day.Day {
	type layer struct {
		d, r int
	}

	readLayers := func() []layer {
		f := input.Open(13)
		scanner := bufio.NewScanner(f)
		layers := make([]layer, 0, 0)
		for scanner.Scan() {
			line := scanner.Text()
			var d, r int
			fmt.Sscanf(line, "%d: %d", &d, &r)
			l := layer{d, r}
			layers = append(layers, l)
		}
		return layers
	}

	part1 := func() interface{} {
		layers := readLayers()
		var penalty int
		for _, layer := range layers {
			if layer.d%((layer.r-1)*2) == 0 {
				penalty += layer.d * layer.r
			}
		}

		return penalty
	}

	part2 := func() interface{} {
		layers := readLayers()
		/*
			sort.Slice(layers, func(i, j int) bool {
				return layers[i].r < layers[j].r
			})
		*/

		pass := func(i int) bool {
			for _, layer := range layers {
				if (layer.d+i)%((layer.r-1)*2) == 0 {
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
