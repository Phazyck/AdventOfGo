package days

import (
	"fmt"

	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
)

type generator func() int

func newGenerator(value, factor, divisor int) generator {
	return func() int {
		value = (value * factor) % divisor
		return value
	}
}

func newPickyGenerator(value, factor, divisor, multipleOf int) generator {
	gen := newGenerator(value, factor, divisor)
	return func() int {
		for {
			v := gen()
			if v%multipleOf == 0 {
				return v
			}
		}
	}
}

func judge(generatorA, generatorB generator, iterations int) int {
	total := 0
	for i := 0; i < iterations; i++ {
		a := generatorA()
		b := generatorB()
		if a&0xFFFF == b&0xFFFF {
			total++
		}
	}
	return total
}

// Day15 is the 15th day in Advent of Code.
// (Actually, it's a template for implementing days.)
func Day15() *day.Day {
	readInput := func() (int, int) {
		lines := input.ReadLines(15, 2)

		var ignore string
		var startA, startB int
		fmt.Sscanf(lines[0], "Generator %s starts with %d", &ignore, &startA)
		fmt.Sscanf(lines[1], "Generator %s starts with %d", &ignore, &startB)
		return startA, startB
	}

	part1 := func() interface{} {
		startA, startB := readInput()
		generatorA := newGenerator(startA, 16807, 2147483647)
		generatorB := newGenerator(startB, 48271, 2147483647)
		total := judge(generatorA, generatorB, 40000000)
		return total
	}

	part2 := func() interface{} {
		startA, startB := readInput()
		generatorA := newPickyGenerator(startA, 16807, 2147483647, 4)
		generatorB := newPickyGenerator(startB, 48271, 2147483647, 8)
		total := judge(generatorA, generatorB, 5000000)
		return total
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.New(15, "Dueling Generators", solve)
}
