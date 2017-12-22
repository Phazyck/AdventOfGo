package main

import (
	"encoding/hex"
	"strconv"
	"strings"
)

func day10() *day {

	parseCommaSeparated := func(input string) []byte {
		parts := strings.Split(input, ",")
		bytes := make([]byte, 0)
		for _, part := range parts {
			i, err := strconv.ParseInt(part, 10, 9)
			check(err)
			b := byte(i)
			bytes = append(bytes, b)
		}
		return bytes
	}

	parseASCII := func(input string) []byte {
		ascii := []byte(input)
		ascii = append(ascii, []byte{17, 31, 73, 47, 23}...)
		return ascii
	}

	makeList := func(size int) []byte {
		list := make([]byte, size, size)
		for i := 0; i < size; i++ {
			list[i] = byte(i)
		}
		return list
	}

	part1 := func() interface{} {
		list := makeList(256)
		var pos, skip byte
		str := readInputLine(10)
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
		list := makeList(256)

		str := readInputLine(10)
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

		hash := hex.EncodeToString(dense)

		return hash
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return &day{10, "Knot Hash", solve}
}
