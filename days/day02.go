package days

import (
	"encoding/csv"
	"io"
	"math"
	"os"
	"sort"

	"github.com/phazyck/adventofgo/assert"
	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
	"github.com/phazyck/adventofgo/parse"
)

// Day02 is the 2nd day in Advent of Code.
func Day02() *day.Day {

	sumRows := func(f *os.File, computeRow func([]int) int) (sum int) {
		r := csv.NewReader(f)
		r.Comma = '\t'

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			assert.IsNil(err)

			row := make([]int, len(record))

			for idx, value := range record {
				i := parse.AssertInt(value)
				row[idx] = i
			}

			sum += computeRow(row)
		}

		return
	}

	part1 := func() interface{} {
		checkRow := func(row []int) int {
			min, max := math.MaxInt64, math.MinInt64

			for _, u := range row {
				if u < min {
					min = u
				}

				if u > max {
					max = u
				}
			}

			return max - min
		}

		f := input.Open(2)
		s := sumRows(f, checkRow)
		return s
	}

	part2 := func() interface{} {
		checkRow := func(row []int) int {
			sort.Ints(row)
			for i, den := range row {
				for j := i + 1; j < len(row); j++ {
					num := row[j]

					div := num / den
					if div*den == num {
						return div
					}
				}
			}

			panic("This shouldn't happen!")
		}

		f := input.Open(2)
		s := sumRows(f, checkRow)
		return s
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.New(2, "Corruption parse.Checksum", solve)
}
