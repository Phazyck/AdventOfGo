package days

import (
	"bufio"
	"fmt"

	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
	"github.com/phazyck/adventofgo/parse"
)

// Day06 is the 6th day in Advent of Code.
func Day06() *day.Day {

	solve := func() (interface{}, interface{}) {
		f := input.Open(6)
		scanner := bufio.NewScanner(f)
		splitTabs := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			// Return nothing if at end of file and no data passed
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}

			for i := 0; i < len(data); i++ {
				b := data[i]
				if b == '\t' || b == '\n' {
					return i + 1, data[0:i], nil
				}
			}

			if atEOF {
				return len(data), data, nil
			}

			return
		}
		scanner.Split(splitTabs)

		banks := make([]int, 0)
		for scanner.Scan() {
			word := scanner.Text()
			i := parse.AssertInt(word)
			banks = append(banks, i)
		}

		realloc := func(memory []int) {
			var idx, val int
			for i, v := range memory {
				if v > val {
					idx, val = i, v
				}
			}

			memory[idx] = 0
			for offset := 1; offset <= val; offset++ {
				i := (idx + offset) % len(memory)
				memory[i]++
			}
		}

		seen := make(map[string]int)

		steps := 0
		var size int
		for {
			s := fmt.Sprint(banks)
			if step, ok := seen[s]; ok {
				size = steps - step
				break
			}
			seen[s] = steps
			realloc(banks)
			steps++
		}

		return steps, size
	}

	return day.New(6, "Memory Reallocation", solve)
}
