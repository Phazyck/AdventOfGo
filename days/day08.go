package days

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/Phazyck/AdventOfGo/common"
	"github.com/Phazyck/AdventOfGo/day"
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

		f := common.OpenInput(8)

		scanner := bufio.NewScanner(f)
		allTimeMax := math.MinInt64

		for scanner.Scan() {
			line := scanner.Text()
			segments := strings.Split(line, " ")

			register := segments[4]
			operator := segments[5]
			value := common.ParseInt(segments[6])

			if !eval(registers[register], operator, value) {
				continue
			}

			register = segments[0]
			operator = segments[1]
			value = common.ParseInt(segments[2])

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

	return day.NewDay(8, "I Heard You Like Registers", solve)
}
