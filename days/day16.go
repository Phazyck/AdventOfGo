package days

import (
	"fmt"
	"strings"

	"github.com/phazyck/adventofgo/day"
	"github.com/phazyck/adventofgo/input"
	"github.com/phazyck/adventofgo/parse"
)

type programs struct {
	dat []byte
	pos int
	len int
}

func newPrograms(chars int) *programs {
	dat := make([]byte, chars, chars)
	for i := 0; i < chars; i++ {
		dat[i] = 'a' + byte(i)
	}

	pos := 0
	len := chars
	return &programs{dat, pos, len}
}

func (ps *programs) String() string {
	pos := ps.pos
	len := ps.len
	dat := ps.dat

	front, back := dat[pos:len], dat[0:pos]
	dat = append(front, back...)
	return string(dat)
}

type move interface {
	apply(*programs)
}

type moveSpin int

func (m moveSpin) apply(ps *programs) {
	pos := ps.pos
	len := ps.len
	pos = (pos - int(m) + len) % len
	ps.pos = pos
}

func (m moveSpin) String() string {
	return fmt.Sprintf("s%d", int(m))
}

type moveExchange struct {
	a int
	b int
}

func (m moveExchange) apply(ps *programs) {
	dat := ps.dat
	pos := ps.pos
	len := ps.len
	a := (pos + m.a) % len
	b := (pos + m.b) % len
	dat[a], dat[b] = dat[b], dat[a]
}

func (m moveExchange) String() string {
	return fmt.Sprintf("x%d/%d", m.a, m.b)
}

type movePartner struct {
	a byte
	b byte
}

func (m movePartner) apply(ps *programs) {
	dat := ps.dat
	a := -1
	b := -1

	for i, v := range dat {
		if v == m.a {
			a = i
			if b != -1 {
				break
			}
		} else if v == m.b {
			b = i
			if a != -1 {
				break
			}
		}
	}

	dat[a], dat[b] = dat[b], dat[a]
}

func (m movePartner) String() string {
	return fmt.Sprintf("p%s/%s", string(m.a), string(m.b))
}

func parseMove(str string) move {
	m, dat := str[0], str[1:]
	switch m {
	case 's': // spin
		{
			var n int
			fmt.Sscanf(dat, "%d", &n)
			//fmt.Printf("Spin %d\n", n)
			return moveSpin(n)
		}
	case 'x': // exchange
		{
			var a, b int
			fmt.Sscanf(dat, "%d/%d", &a, &b)

			//fmt.Printf("Exchange %d with %d\n", a, b)
			return moveExchange{a, b}

		}
	case 'p': // partner
		{
			parts := strings.Split(dat, "/")
			a := parse.AssertChar(parts[0])
			b := parse.AssertChar(parts[1])

			//fmt.Printf("Partner %s with %s\n", string(a), string(b))
			return movePartner{a, b}
		}
	default:
		{
			message := fmt.Sprintf("Unknown move '%s'", string(m))
			panic(message)
		}
	}

}

// Day16 is the 16th day in Advent of Code.
func Day16() *day.Day {

	loadDance := func() []move {
		line := input.ReadLine(16)
		strs := strings.Split(line, ",")

		dance := make([]move, len(strs), len(strs))
		for i, str := range strs {
			dance[i] = parseMove(str)
		}

		return dance
	}

	solve := func() (interface{}, interface{}) {
		dance := loadDance()
		prog := newPrograms(16)
		init := prog.String()

		permutations := []string{init}

		for {
			for _, move := range dance {
				move.apply(prog)
			}

			str := prog.String()

			if str == init {
				break
			}
			permutations = append(permutations, str)
		}

		first := permutations[1%len(permutations)]
		final := permutations[1000000000%len(permutations)]

		return first, final
	}

	return day.New(16, "Permutation Promenade", solve)
}
