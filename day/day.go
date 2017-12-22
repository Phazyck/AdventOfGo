package day

import "fmt"

// Solver is a function that returns the two solutions for a given day.
type Solver func() (interface{}, interface{})

// Day contains the information and implementation of solutions to a day in Advent of Code.
type Day struct {
	number int
	title  string
	solve  Solver
}

// New returns a new day with the given values assigned.
func New(number int, title string, solve Solver) *Day {
	return &Day{number, title, solve}
}

// Print pretty-prints the information for a given day on multiple lines.
func (d *Day) Print() {
	part1, part2 := d.solve()
	fmt.Printf("--- Day %02d: %s ---\n", d.number, d.title)
	fmt.Println(part1)
	fmt.Println()
	fmt.Println("--- Part Two ---")
	fmt.Println(part2)
	fmt.Println()
}

// PrintShort pretty-prints the information for a given day on one line.
func (d *Day) PrintShort() {
	part1, part2 := d.solve()
	fmt.Printf("[%02d.1] %v\n", d.number, part1)
	fmt.Printf("[%02d.2] %v\n", d.number, part2)
}
