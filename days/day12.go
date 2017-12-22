package days

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/Phazyck/AdventOfGo/common"
	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/stack"
)

// Day12 is the 12th day in Advent of Code.
func Day12() *day.Day {

	readPipes := func(f *os.File) map[int][]int {
		pipes := make(map[int][]int)
		scanner := bufio.NewScanner(f)

		reL := regexp.MustCompile("(.*) <->")
		reR := regexp.MustCompile("<-> (.*)")

		for scanner.Scan() {
			line := scanner.Text()

			strL := reL.FindStringSubmatch(line)[1]
			strR := reR.FindStringSubmatch(line)[1]

			l := common.ParseInt(strL)
			strsR := strings.Split(strR, ", ")
			rCount := len(strsR)
			rs := make([]int, rCount, rCount)
			for idx, str := range strsR {
				rs[idx] = common.ParseInt(str)
			}

			pipes[l] = rs
		}

		return pipes
	}

	findGroup := func(pipes map[int][]int, id int) map[int]bool {
		group := make(map[int]bool)

		next := stack.NewStack(id)

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
		f := common.OpenInput(12)
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

	return day.NewDay(12, "Digital Plumber", solve)
}
