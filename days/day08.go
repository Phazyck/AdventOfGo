package days

import (
	"bufio"
	"fmt"
	"math"

	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
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

			var rMod, rTst string // (r)egisters
			var oMod, oTst string // (o)perators
			var vMod, vTst int    // (v)alues
			fmt.Sscanf(line, "%s %s %d if %s %s %d", &rMod, &oMod, &vMod, &rTst, &oTst, &vTst)

			if !eval(registers[rTst], oTst, vTst) {
				continue
			}

			value := registers[rMod]
			if oMod == "inc" {
				value += vMod
			} else {
				value -= vMod
			}

			registers[rMod] = value

			if value > allTimeMax {
				allTimeMax = value
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
