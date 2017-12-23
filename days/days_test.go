package days

import (
	"testing"

	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/test"
)

func testDay(t *testing.T, d *day.Day, expected1, expected2 interface{}) {
	actual1, actual2 := d.Solutions()
	test.AssertEqual(t, actual1, expected1)
	test.AssertEqual(t, actual2, expected2)
}

// TestDay01Solutions tests that the solutions for day 01 are correct.
func TestDay01Solutions(t *testing.T) {
	testDay(t, Day01(), 1175, 1166)
}

// TestDay02Solutions tests that the solutions for day 02 are correct.
func TestDay02Solutions(t *testing.T) {
	testDay(t, Day02(), 47136, 250)
}

// TestDay03Solutions tests that the solutions for day 03 are correct.
func TestDay03Solutions(t *testing.T) {
	testDay(t, Day03(), 480, 349975)
}

// TestDay04Solutions tests that the solutions for day 04 are correct.
func TestDay04Solutions(t *testing.T) {
	testDay(t, Day04(), 325, 119)
}

// TestDay05Solutions tests that the solutions for day 05 are correct.
func TestDay05Solutions(t *testing.T) {
	testDay(t, Day05(), 376976, 29227751)
}

// TestDay06Solutions tests that the solutions for day 06 are correct.
func TestDay06Solutions(t *testing.T) {
	testDay(t, Day06(), 5042, 1086)
}

// TestDay07Solutions tests that the solutions for day 07 are correct.
func TestDay07Solutions(t *testing.T) {
	testDay(t, Day07(), "vmpywg", 1674)
}

// TestDay08Solutions tests that the solutions for day 08 are correct.
func TestDay08Solutions(t *testing.T) {
	testDay(t, Day08(), 4163, 5347)
}

// TestDay09Solutions tests that the solutions for day 09 are correct.
func TestDay09Solutions(t *testing.T) {
	testDay(t, Day09(), 14204, 6622)
}

// TestDay10Solutions tests that the solutions for day 10 are correct.
func TestDay10Solutions(t *testing.T) {
	testDay(t, Day10(), 15990, "90adb097dd55dea8305c900372258ac6")
}

func testDay11Example(t *testing.T, input string, expected interface{}) {
	actual, _ := solveInput(input)
	test.AssertEqual(t, actual, expected)
}

// TestDay11Examples tests that solveInput works for the examples of day 11
func TestDay11Examples(t *testing.T) {
	testDay11Example(t, "ne,ne,ne", 3)
	testDay11Example(t, "ne,ne,sw,sw", 0)
	testDay11Example(t, "ne,ne,s,s", 2)
	testDay11Example(t, "se,sw,se,sw,sw", 3)
}

// TestDay11Solutions tests that the solutions for day 11 are correct.
func TestDay11Solutions(t *testing.T) {
	testDay(t, Day11(), 675, 1424)
}

// TestDay12Solutions tests that the solutions for day 12 are correct.
func TestDay12Solutions(t *testing.T) {
	testDay(t, Day12(), 283, 195)
}

// TestDay13Solutions tests that the solutions for day 13 are correct.
func TestDay13Solutions(t *testing.T) {
	testDay(t, Day13(), 1504, 3823370)
}

// TestDay14Solutions tests that the solutions for day 14 are correct.
func TestDay14Solutions(t *testing.T) {
	used, regions := solveDay14("flqrgnkx")
	test.AssertEqual(t, used, 8108)
	test.AssertEqual(t, regions, 1242)

	testDay(t, Day14(), 8148, 1180)
}
