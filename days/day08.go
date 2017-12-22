package days

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/input"
	"github.com/Phazyck/AdventOfGo/parse"
)

// Day08 is the 8th day in Advent of Code.
func Day08() *day.Day {
	/*
		type condition struct {
			register string
			comparison string
			value int
		}

		type instruction struct {
			register string
			operation string
			value int
			condition condition
		}
	*/

	eval := func(a int, op string, b int) bool {
		if op == "==" {
			return a == b
		}

		if op == "!=" {
			return a != b
		}

		if op == ">" {
			return a > b
		}

		if op == "<" {
			return a < b
		}

		if op == "<=" {
			return a <= b
		}

		if op == ">=" {
			return a >= b
		}

		panic(fmt.Sprintf("Unknown operator '%s'", op))
	}

	solve := func() (interface{}, interface{}) {
		registers := make(map[string]int)

		f := input.Open(8)

		scanner := bufio.NewScanner(f)
		allTimeMax := math.MinInt64

		for scanner.Scan() {
			line := scanner.Text()
			segments := strings.Split(line, " ")

			register := segments[4]
			operator := segments[5]
			value := parse.AssertInt(segments[6])

			if !eval(registers[register], operator, value) {
				continue
			}

			register = segments[0]
			operator = segments[1]
			value = parse.AssertInt(segments[2])

			if operator == "inc" {
				registers[register] += value
			} else {
				registers[register] -= value
			}

			if registers[register] > allTimeMax {
				allTimeMax = registers[register]
			}
		}

		max := math.MinInt64

		for _, val := range registers {
			if val > max {
				max = val
			}
		}

		return max, allTimeMax
	}

	return day.New(8, "I Heard You Like Registers", solve)
}
