package days

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
)

// Day19 is the 19th day in Advent of Code.
func Day19() *day.Day {

	solve := func() (interface{}, interface{}) {
		f := input.Open(19)
		scanner := bufio.NewScanner(f)

		maze := make([]string, 0, 0)
		for scanner.Scan() {
			maze = append(maze, scanner.Text())
		}

		x, y := strings.IndexRune(maze[0], '|'), 0
		dx, dy := 0, 1

		var visited bytes.Buffer

		for steps := 1; ; steps++ {
			x, y = x+dx, y+dy
			c := maze[y][x]

			switch c {
			case ' ':
				return visited.String(), steps
			case '+':
				dx, dy = dy, dx
				if maze[y+dy][x+dx] == ' ' {
					dx, dy = -dx, -dy
				}
			case '-': // no-op
			case '|': // no-op
			default:
				fmt.Fprint(&visited, string(c))
			}
		}
	}

	return day.New(19, "A Series of Tubes", solve)
}
