package days

import (
	"bufio"
	"os"
	"strings"

	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/input"
	"github.com/Phazyck/AdventOfGo/parse"
	"github.com/Phazyck/AdventOfGo/stack"
)

// Day12 is the 12th day in Advent of Code.
func Day12() *day.Day {

	readPipes := func(f *os.File) map[int][]int {
		pipes := make(map[int][]int)
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()

			parts := strings.Split(line, " <-> ")
			l := parse.AssertInt(parts[0])
			rs := parse.AssertInts(parts[1], ", ")

			pipes[l] = rs
		}

		return pipes
	}

	findGroup := func(pipes map[int][]int, id int) map[int]bool {
		group := make(map[int]bool)

		next := stack.New(id)

		for !next.Empty() {
			id := next.Pop()

			if group[id] {
				continue
			}

			group[id] = true

			values := pipes[id]
			next.Push(values...)
			delete(pipes, id)
		}

		return group
	}

	solve := func() (interface{}, interface{}) {
		f := input.Open(12)
		pipes := readPipes(f)
		group0 := findGroup(pipes, 0)
		groupCount := 1

		for len(pipes) > 0 {
			var id int
			for k := range pipes {
				id = k
				break
			}
			findGroup(pipes, id)
			groupCount++
		}

		return len(group0), groupCount
	}

	return day.New(12, "Digital Plumber", solve)
}
