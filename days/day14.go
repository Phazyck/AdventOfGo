package days

import (
	"fmt"

	"github.com/Phazyck/AdventOfGo/day"
)

func fillGroup(memory [][]byte, i, j int, id byte) {
	if i < 0 || i >= len(memory) {
		return
	}

	row := memory[i]

	if j < 0 || j >= len(row) {
		return
	}

	cell := row[j]

	if cell != '#' {
		return
	}

	row[j] = id

	fillGroup(memory, i+1, j, id)
	fillGroup(memory, i-1, j, id)
	fillGroup(memory, i, j+1, id)
	fillGroup(memory, i, j-1, id)
}

func print(memory [][]byte) {
	for i, row := range memory {
		fmt.Printf("%3d %s\n", i, string(row))
	}
}

func solveDay14(key string) (interface{}, interface{}) {
	sz := 128
	memory := make([][]byte, sz, sz)

	used := 0
	for i := 0; i < sz; i++ {
		input := fmt.Sprintf("%s-%d", key, i)
		hash := knotHash(input)
		j := 0

		row := make([]byte, sz, sz)

		for _, b := range hash {
			bin := fmt.Sprintf("%08b", b)
			for _, bit := range bin {
				if bit == '1' {
					row[j] = '#'
					used++
				} else {
					row[j] = '.'
				}
				j++
			}
		}

		memory[i] = row
	}

	var groups int
	var groupID byte = '1'

	for i, row := range memory {
		for j, cell := range row {
			if cell == '#' {
				groups++
				fillGroup(memory, i, j, groupID)
				groupID++
				if groupID == 'z'+1 {
					groupID = 'A'
				} else if groupID == 'Z'+1 {
					groupID = '1'
				} else if groupID == '9'+1 {
					groupID = 'a'
				}
			}
		}
	}

	print(memory)
	fmt.Println()

	return used, groups
}

// Day14 is the 14th day in Advent of Code.
func Day14() *day.Day {
	solve := func() (interface{}, interface{}) {
		return solveDay14("vbqugkhl")
	}

	return day.New(14, "Disk Defragmentation", solve)
}
