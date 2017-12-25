package days

import (
	"bytes"
	"fmt"

	"github.com/Phazyck/AdventOfGo/day"
)

type node struct {
	val int
	nxt *node
}

type spinlock struct {
	nds []node
	cur *node
	stp int
	siz int
}

func newSpinlock(stp, cap int) *spinlock {
	nds := make([]node, cap, cap)
	for i := 0; i < cap; i++ {
		nds[i].val = i
	}
	cur := &nds[0]
	return &spinlock{nds, cur, stp, 1}
}

func (sl *spinlock) advanceTo(size int) {
	for sl.siz < size {
		nod := sl.cur
		for i := 0; i < sl.stp; i++ {
			nod = nod.nxt
			if nod == nil {
				nod = &sl.nds[0]
			}
		}

		new := &sl.nds[sl.siz]
		new.nxt = nod.nxt
		nod.nxt = new

		sl.cur = new

		sl.siz++
	}
}

func (sl *spinlock) String() string {
	var b bytes.Buffer
	appendf := func(format string, a ...interface{}) {
		fmt.Fprintf(&b, format, a...)
	}

	nod := &sl.nds[0]
	cur := sl.cur
	val := nod.val
	if nod == cur {
		appendf("(%d)", val)

	} else {
		appendf("%d", val)
	}

	for i := 1; i < sl.siz; i++ {
		nod = nod.nxt
		val = nod.val
		switch nod {
		case cur:
			appendf(" (%d)", val)
		case cur.nxt:
			appendf(" %d", val)
		default:
			appendf("  %d", val)
		}
	}

	return b.String()
}

// Day17 is the 17th day in Advent of Code.
func Day17() *day.Day {

	solve := func() (interface{}, interface{}) {
		stp := 343
		sl := newSpinlock(stp, 2018)
		sl.advanceTo(2018)
		pt1 := sl.nds[2017].nxt.val

		idx := 0
		pt2 := 0
		for siz := 1; siz < 50000000; siz++ {
			idx = ((idx + stp) % siz) + 1
			if idx == 1 {
				pt2 = siz
			}
		}

		return pt1, pt2
	}

	return day.New(17, "Spinlock", solve)
}
