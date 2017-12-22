package main

func day00() *day {

	part1 := func() interface{} {
		return -1
	}

	part2 := func() interface{} {
		return -1
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return &day{00, "Title", solve}
}
