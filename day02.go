package main

import (
	"encoding/csv"
	"io"
	"math"
	"os"
	"sort"
)

func day2() *day {

	sumRows := func(f *os.File, computeRow func([]int) int) (sum int) {
		r := csv.NewReader(f)
		r.Comma = '\t'

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			check(err)

			row := make([]int, len(record))

			for idx, value := range record {
				i := parseInt(value)
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

		f := openInput(2)
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

		f := openInput(2)
		s := sumRows(f, checkRow)
		return s
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return &day{2, "Corruption Checksum", solve}
}
