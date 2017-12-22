package main

import "fmt"

type day struct {
	number int
	title  string
	solve  func() (interface{}, interface{})
}

func (d *day) print() {
	part1, part2 := d.solve()
	fmt.Printf("--- Day %02d: %s ---\n", d.number, d.title)
	fmt.Println(part1)
	fmt.Println()
	fmt.Println("--- Part Two ---")
	fmt.Println(part2)
	fmt.Println()
}

func (d *day) printShort() {
	part1, part2 := d.solve()
	fmt.Printf("[%02d.1] %v\n", d.number, part1)
	fmt.Printf("[%02d.2] %v\n", d.number, part2)
}
