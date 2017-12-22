package main

func printAll(days []*day) {
	for _, day := range days {
		day.printShort()
	}
}

func printLatest(days []*day) {
	l := len(days)
	i := l - 1
	day := days[i]
	day.printShort()
}

func main() {
	days := []*day{
		day01(),
		day02(),
		day03(),
		day04(),
		day05(),
		day06(),
		day07(),
		day08(),
		day09(),
		day10(),
		day11(),
		day12(),
	}

	//printAll(days)
	printLatest(days)
}
