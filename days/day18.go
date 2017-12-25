package days

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/input"
	"github.com/Phazyck/AdventOfGo/stack"
)

type duet struct {
	rgs map[string]int
	snd int
}

type duet2 struct {
	rgs map[string]int
	snd *stack.Stack
	rcv *stack.Stack
}

/*
func newQueue() (rcv <-chan int, snd chan<- int) {
	r := make(chan int)
	s := make(chan int)

	transfer := func() {
		buf := make([]int, 0, 0)

		for {
			if len(buf) == 0 {
				i := <-s
				buf = append(buf, i)
				fmt.Printf("snd %d, len(buf)=%d", i, len(buf))
			} else {
				o := buf[0]
				select {
				case r <- o:
					buf = buf[1:]
					fmt.Printf("rcv %d, len(buf)=%d", o, len(buf))
				case i := <-s:
					buf = append(buf, i)
					fmt.Printf("snd %d, len(buf)=%d", i, len(buf))
				}
			}
		}
	}
	go transfer()

	return r, s
}
*/

func newDuet() *duet {
	rgs := make(map[string]int)
	snd := 0
	return &duet{rgs, snd}
}

func newDuet2() (*duet2, *duet2) {
	rgs0 := map[string]int{
		"p": 0,
	}
	rgs1 := map[string]int{
		"p": 1,
	}

	stk0 := stack.New()
	stk1 := stack.New()

	d0 := &duet2{rgs0, stk1, stk0}
	d1 := &duet2{rgs1, stk0, stk1}
	return d0, d1
}

func (d *duet2) exec(ins []string) (jmp int, blk bool) {
	eval := func(str string) int {
		val, err := strconv.Atoi(str)
		if err == nil {
			return val
		}
		return d.rgs[str]
	}

	var o, x, y string
	o = ins[0]
	x = ins[1]

	if len(ins) > 2 {
		y = ins[2]
	}

	jmp = 1
	switch o {
	case "snd":
		d.snd.Push(eval(x))
	case "set":
		d.rgs[x] = eval(y)
	case "add":
		d.rgs[x] += eval(y)
	case "mul":
		d.rgs[x] *= eval(y)
	case "mod":
		d.rgs[x] %= eval(y)
	case "rcv":
		if d.rcv.Empty() {
			blk = true
			return
		}
		d.rgs[x] = d.rcv.Pop()
	case "jgz":
		if eval(x) > 0 {
			jmp = eval(y)
		}
	}

	return
}

func (d *duet) exec(ins []string) (jmp int, rcv bool) {

	eval := func(str string) int {
		val, err := strconv.Atoi(str)
		if err == nil {
			return val
		}
		return d.rgs[str]
	}

	var o, x, y string
	o = ins[0]
	x = ins[1]

	if len(ins) > 2 {
		y = ins[2]
	}

	jmp = 1
	switch o {
	case "snd":
		d.snd = eval(x)
	case "set":
		d.rgs[x] = eval(y)
	case "add":
		d.rgs[x] += eval(y)
	case "mul":
		d.rgs[x] *= eval(y)
	case "mod":
		d.rgs[x] %= eval(y)
	case "rcv":
		if eval(x) != 0 {
			rcv = true
		}
	case "jgz":
		if eval(x) > 0 {
			jmp = eval(y)
		}
	}

	return
}

func parseInstruction(str string) []string {
	ins := strings.Split(str, " ")
	return ins
}

func parseInstructions(ipt []string) [][]string {
	iss := make([][]string, len(ipt), len(ipt))
	for idx, str := range ipt {
		iss[idx] = parseInstruction(str)
	}
	return iss
}

// Day18 is the 18th day in Advent of Code.
func Day18() *day.Day {

	loadInstructions := func() [][]string {
		f := input.Open(18)
		scanner := bufio.NewScanner(f)

		iss := make([][]string, 0, 0)
		for scanner.Scan() {
			str := scanner.Text()
			ins := parseInstruction(str)
			iss = append(iss, ins)
		}
		return iss
	}

	part1 := func() interface{} {
		iss := loadInstructions()

		d := newDuet()

		isc := 0
		for {
			ins := iss[isc]
			jmp, rcv := d.exec(ins)
			isc += jmp
			if rcv {
				break
			}
		}

		return d.snd
	}

	part2 := func() interface{} {
		iss := loadInstructions()

		prg0, prg1 := newDuet2()
		a, b := prg0, prg1

		ica, icb := 0, 0

		count := 0

		swapped := false

		for {
			ins := iss[ica]

			if a == prg1 && ins[0] == "snd" {
				count++
			}

			jmp, blk := a.exec(ins)

			if blk {
				// program is blocking

				if swapped {
					// already swapped once -> deadlock
					break
				}

				// swap programs
				a, b = b, a
				ica, icb = icb, ica
				swapped = true
			} else {
				ica += jmp
				swapped = false
			}
		}

		return count
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return day.New(18, "Duet", solve)
}
