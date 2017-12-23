package days

import (
	"fmt"
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

// TestDay14Examples tests that solveInput works for the examples of day 14
func TestDay14Examples(t *testing.T) {
	used, regions := solveDay14("flqrgnkx")
	test.AssertEqual(t, used, 8108)
	test.AssertEqual(t, regions, 1242)
}

// TestDay14Solutions tests that the solutions for day 14 are correct.
func TestDay14Solutions(t *testing.T) {
	testDay(t, Day14(), 8148, 1180)
}

// TestDay15Examples tests that solveInput works for the examples of day 15
func TestDay15Examples(t *testing.T) {

	testGenerator := func(gen generator, expectedInts []int, expectedBinaries []string) {
		for i, expectedInt := range expectedInts {
			expectedBinary := expectedBinaries[i]
			actualInt := gen()
			test.AssertEqual(t, actualInt, expectedInt)
			actualBinary := fmt.Sprintf("%032b", actualInt)
			test.AssertEqual(t, actualBinary, expectedBinary)
		}
	}

	newGeneratorA := func() generator {
		return newGenerator(65, 16807, 2147483647)
	}

	newPickyA := func() generator {
		return newPickyGenerator(65, 16807, 2147483647, 4)
	}

	newGeneratorB := func() generator {
		return newGenerator(8921, 48271, 2147483647)
	}

	newPickyB := func() generator {
		return newPickyGenerator(8921, 48271, 2147483647, 8)
	}

	testGenerator(newGeneratorA(), []int{
		1092455,
		1181022009,
		245556042,
		1744312007,
		1352636452,
	}, []string{
		"00000000000100001010101101100111",
		"01000110011001001111011100111001",
		"00001110101000101110001101001010",
		"01100111111110000001011011000111",
		"01010000100111111001100000100100",
	})

	testGenerator(newGeneratorB(), []int{
		430625591,
		1233683848,
		1431495498,
		137874439,
		285222916,
	}, []string{
		"00011001101010101101001100110111",
		"01001001100010001000010110001000",
		"01010101010100101110001101001010",
		"00001000001101111100110000000111",
		"00010001000000000010100000000100",
	})

	testGenerator(newPickyA(), []int{
		1352636452,
		1992081072,
		530830436,
		1980017072,
		740335192,
	}, []string{
		"01010000100111111001100000100100",
		"01110110101111001011111010110000",
		"00011111101000111101010001100100",
		"01110110000001001010100110110000",
		"00101100001000001001111001011000",
	})

	testGenerator(newPickyB(), []int{
		1233683848,
		862516352,
		1159784568,
		1616057672,
		412269392,
	}, []string{
		"01001001100010001000010110001000",
		"00110011011010001111010010000000",
		"01000101001000001110100001111000",
		"01100000010100110001010101001000",
		"00011000100100101011101101010000",
	})

	actualTotal5 := judge(newGeneratorA(), newGeneratorB(), 5)
	test.AssertEqual(t, actualTotal5, 1)

	actualTotal40M := judge(newGeneratorA(), newGeneratorB(), 40000000)
	test.AssertEqual(t, actualTotal40M, 588)

	actualTotalPicky5M := judge(newPickyA(), newPickyB(), 5000000)
	test.AssertEqual(t, actualTotalPicky5M, 309)
}

// TestDay15Solutions tests that the solutions for day 15 are correct.
func TestDay15Solutions(t *testing.T) {
	testDay(t, Day15(), 609, 253)
}
