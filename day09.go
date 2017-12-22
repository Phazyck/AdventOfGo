package main

import "bufio"

func day09() *day {

	countGarbage := func(r *bufio.Reader) int {
		count := 0
		for {
			c, _, err := r.ReadRune()
			check(err)
			switch c {
			case '>':
				return count
			case '!':
				_, _, err = r.ReadRune()
				check(err)
			default:
				count++
			}
		}
	}

	solve := func() (interface{}, interface{}) {
		f := openInput(9)
		r := bufio.NewReader(f)

		total := 0
		score := 0
		garbage := 0
		for {
			c, _, err := r.ReadRune()

			if err != nil {
				break
			}

			switch c {
			case '{':
				score++

			case '}':
				total += score
				score--
			case '!':
				_, _, err = r.ReadRune()
				check(err)
			case '<':
				garbage += countGarbage(r)
			}
		}

		return total, garbage
	}

	return &day{9, "Stream Processing", solve}
}
