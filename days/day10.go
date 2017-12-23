package days

import (
	"encoding/hex"
	"strings"

	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/input"
	"github.com/Phazyck/AdventOfGo/parse"
)

func makeList(size int) []byte {
	list := make([]byte, size, size)
	for i := range list {
		list[i] = byte(i)
	}
	return list
}

func parseASCII(input string) []byte {
	ascii := []byte(input)
	ascii = append(ascii, []byte{17, 31, 73, 47, 23}...)
	return ascii
}

func knotHash(str string) []byte {
	list := makeList(256)
	input := parseASCII(str)
	var pos, skip byte

	for round := 0; round < 64; round++ {
		for _, length := range input {
			for i := byte(0); i < length/2; i++ {
				j := length - 1 - i
				pi := pos + i
				pj := pos + j
				list[pi], list[pj] = list[pj], list[pi]
			}

			pos += length + skip
			skip++
		}
	}

	densify := func(block []byte) byte {
		value := block[0]
		for i := 1; i < len(block); i++ {
			value = value ^ block[i]
		}
		return value
	}

	dense := make([]byte, 16, 16)

	for i := 0; i < 16; i++ {
		from := i * 16
		to := from + 16
		block := list[from:to]
		dense[i] = densify(block)
	}

	return dense
}

// Day10 is the 10th day in Advent of Code.
func Day10() *day.Day {

	parseCommaSeparated := func(input string) []byte {
		parts := strings.Split(input, ",")
		bytes := make([]byte, 0)
		for _, part := range parts {
			i := parse.AssertByte(part)
			b := byte(i)
			bytes = append(bytes, b)
		}
		return bytes
	}

	part1 := func() interface{} {
		list := makeList(256)
		var pos, skip byte
		str := input.ReadLine(10)
		input := parseCommaSeparated(str)

		for _, length := range input {
			for i := byte(0); i < length/2; i++ {
				j := length - 1 - i
				pi := pos + i
				pj := pos + j
				list[pi], list[pj] = list[pj], list[pi]
			}

			pos += length + skip
			skip++
		}

		return int(list[0]) * int(list[1])
	}

	part2 := func() interface{} {
		str := input.ReadLine(10)
		hash := knotHash(str)
		return hex.EncodeToString(hash)
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.New(10, "Knot Hash", solve)
}
