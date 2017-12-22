package days

import "github.com/Phazyck/AdventOfGo/day"

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func manhattan(x, y int) int {
	return abs(x) + abs(y)
}

type spiral struct {
	squares map[int]map[int]int
}

func (s *spiral) set(x, y, value int) {
	ys, ok := s.squares[x]
	if !ok {
		ys = make(map[int]int)
		s.squares[x] = ys
	}
	ys[y] = value
}

func (s *spiral) get(x, y int) (value int, ok bool) {
	ys, ok := s.squares[x]
	if !ok {
		return
	}
	value, ok = ys[y]
	return
}

// Day03 is the 3rd day in Advent of Code.
func Day03() *day.Day {

	/*
		steps := func(addr int) int {
			if addr == 1 {
				return 0
			}

			idx := addr - 1

			maxSteps := int(math.Sqrt(float64(idx))) + 1
			minSteps := maxSteps / 2

			width := 1 + maxSteps

			offset := idx%maxSteps - width/2
			if offset < 0 {
				offset = -offset
			}

			return minSteps + offset
		}
	*/

	type spiralBuilder func(*spiral, int, int, int) (exit bool)

	buildSpiral := func(build spiralBuilder) *spiral {
		spiral := &spiral{make(map[int]map[int]int)}

		x, y := 0, 0
		dx, dy := 1, 0

		l1, l2 := 1, 1

		i := 1
		for {
			for step := 0; step < l1; step++ {
				if build(spiral, x, y, i) {
					return spiral
				}
				i++
				x += dx
				y += dy
			}

			dx, dy = -dy, dx
			l1, l2 = l2, (l1 + 1)
		}
	}

	part1 := func() interface{} {
		input := 347991
		var solution int
		builder := func(s *spiral, x, y, i int) bool {
			s.set(x, y, i)
			if i == input {
				solution = manhattan(x, y)
				return true
			}
			return false
		}
		buildSpiral(builder)
		return solution
	}

	part2 := func() interface{} {
		input := 347991
		var solution int
		build := func(s *spiral, x, y, i int) bool {
			var value int
			if i == 1 {
				value = 1
			} else {
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if dx == 0 && dy == 0 {
							continue
						}
						v, ok := s.get(x+dx, y+dy)
						if ok {
							value += v
						}
					}
				}
			}

			//fmt.Printf("[%2d] @ (%2d,%2d) = %d", i, x, y, sum)
			s.set(x, y, value)
			if value > input {
				solution = value
				return true
			}
			return false
		}
		buildSpiral(build)
		return solution
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.New(3, "Spiral Memory", solve)
}
