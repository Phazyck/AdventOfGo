package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func day12() *day {

	readPipes := func(f *os.File) map[int][]int {
		pipes := make(map[int][]int)
		scanner := bufio.NewScanner(f)

		reL := regexp.MustCompile("(.*) <->")
		reR := regexp.MustCompile("<-> (.*)")

		for scanner.Scan() {
			line := scanner.Text()

			strL := reL.FindStringSubmatch(line)[1]
			strR := reR.FindStringSubmatch(line)[1]

			l := parseInt(strL)
			strsR := strings.Split(strR, ", ")
			rCount := len(strsR)
			rs := make([]int, rCount, rCount)
			for idx, str := range strsR {
				rs[idx] = parseInt(str)
			}

			pipes[l] = rs
		}

		return pipes
	}

	findGroup := func(pipes map[int][]int, id int) map[int]bool {
		group := make(map[int]bool)

		next := stack{[]int{id}}

		for !next.empty() {
			id := next.pop()

			if group[id] {
				continue
			}

			group[id] = true

			values := pipes[id]
			next.push(values...)
			delete(pipes, id)
		}

		return group
	}

	solve := func() (interface{}, interface{}) {
		f := openInput(12)
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

	return &day{12, "Digital Plumber", solve}
}
