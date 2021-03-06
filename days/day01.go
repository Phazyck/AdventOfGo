package days

import (
	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
)

func toDigits(bytes []byte) []byte {
	digits := make([]byte, 0)
	for _, b := range bytes {
		zero := byte('0')
		d := b - zero
		if d >= 0 && d <= 9 {
			digits = append(digits, d)
		}
	}
	return digits
}

// Day01 is the 1st day in Advent of Code.
func Day01() *day.Day {

	solveCaptcha := func(captcha []byte, offset int) (solution int) {
		l := len(captcha)

		for idx := 0; idx < l; idx++ {
			d1 := captcha[idx]
			j := (idx + offset) % l
			d2 := captcha[j]

			if d1 == d2 {
				solution += int(d1)
			}
		}

		return
	}

	solve := func() (interface{}, interface{}) {
		bytes := input.Read(1)
		captcha := toDigits(bytes)
		part1 := solveCaptcha(captcha, 1)
		part2 := solveCaptcha(captcha, len(captcha)/2)
		return part1, part2
	}

	return day.New(1, "Inverse Captcha", solve)
}
