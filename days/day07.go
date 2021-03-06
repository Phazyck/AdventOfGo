package days

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
)

type tower struct {
	below  map[string]string
	above  map[string][]string
	weight map[string]int
	root   string
}

func (t *tower) add(name string, weight int, above []string) {
	t.weight[name] = weight

	if above != nil {
		t.above[name] = above
		for _, n := range above {
			t.below[n] = name
		}
	}
}

func (t *tower) find(name string) (weight int, found bool) {
	names, ok := t.above[name]
	weight = t.weight[name]
	if !ok {
		return weight, false
	}

	weights := make(map[int][]string)

	for _, n := range names {
		w, found := t.find(n)
		if found {
			return w, true
		}

		weight += w

		if v, ok := weights[w]; !ok {
			weights[w] = []string{n}
		} else {
			weights[w] = append(v, n)
		}
	}

	if len(names) < 3 {
		return weight, false
	}

	var valMin, valMax []string
	var keyMin, keyMax int

	for key, val := range weights {
		if key > keyMax || valMax == nil {
			keyMax, valMax = key, val
		}

		if key < keyMin || valMin == nil {
			keyMin, valMin = key, val
		}
	}

	if keyMax != keyMin {
		var name string
		var actual, target int
		if len(valMax) == 1 {
			name = valMax[0]
			actual = keyMax
			target = keyMin
		} else {
			name = valMin[0]
			actual = keyMin
			target = keyMax
		}
		change := target - actual
		current := t.weight[name]
		return current + change, true
	}

	return weight, false
}

// Day07 is the 7th day in Advent of Code.
func Day07() *day.Day {

	build := func() *tower {
		f := input.Open(7)

		scanner := bufio.NewScanner(f)

		tower := &tower{
			make(map[string]string),
			make(map[string][]string),
			make(map[string]int),
			"",
		}

		candidates := make(map[string]bool)

		for scanner.Scan() {
			line := scanner.Text()

			var name string
			var weight int
			fmt.Sscanf(line, "%s (%d)", &name, &weight)

			candidates[name] = true

			var above []string
			arrow := " -> "
			i := strings.Index(line, arrow)
			if i >= 0 {
				values := line[i+len(arrow):]
				above = strings.Split(values, ", ")
				for _, n := range above {
					delete(candidates, n)
				}
			}

			tower.add(name, weight, above)
		}

		for k := range tower.below {
			delete(candidates, k)
		}

		if len(candidates) != 1 {
			panic("This is unexptected...")
		}

		for k := range candidates {
			tower.root = k
		}

		return tower
	}

	solve := func() (interface{}, interface{}) {
		tower := build()
		part1 := tower.root
		part2, _ := tower.find(tower.root)
		return part1, part2
	}

	return day.New(7, "Recursive Circus", solve)
}
