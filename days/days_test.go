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

// TestDay11Example1 tests that solveInput works for the 1st example of day 11
func TestDay11Example1(t *testing.T) {
	result, _ := solveInput("ne,ne,ne")
	test.AssertEqual(t, result, 3)
}

// TestDay11Example2 tests that solveInput works for the 2nd example of day 11
func TestDay11Example2(t *testing.T) {
	result, _ := solveInput("ne,ne,sw,sw")
	test.AssertEqual(t, result, 0)
}

// TestDay11Example3 tests that solveInput works for the 3rd example of day 11
func TestDay11Example3(t *testing.T) {
	result, _ := solveInput("ne,ne,s,s")
	test.AssertEqual(t, result, 2)
}

// TestDay11Example4 tests that solveInput works for the 4th example of day 11
func TestDay11Example4(t *testing.T) {
	result, _ := solveInput("se,sw,se,sw,sw")
	test.AssertEqual(t, result, 3)
}
