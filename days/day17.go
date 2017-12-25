package days

import (
	"fmt"
	"bytes"

	"github.com/Phazyck/AdventOfGo/day"
)

type spinlock struct {
	buf []int
	pos int
	stp int
	cap int
}

func newSpinlock(stp, cap int) *spinlock {
	buf := make([]int, 1, cap)
	pos := 0
	return &spinlock{buf,pos,stp,cap}
}

func (sl *spinlock) advanceTo(size int) {
	for len(sl.buf) < size {
		i := len(sl.buf)
		
		sl.pos = ((sl.pos + sl.stp) % i) + 1
	
		// insert at pos
		sl.buf = append(sl.buf, 0)
		dst := sl.buf[sl.pos+1:]
		src := sl.buf[sl.pos:]
		copy(dst, src)
		sl.buf[sl.pos] = i
	}
}

func (sl *spinlock) size() int {
	return len(sl.buf)
}

func (sl *spinlock) get(idx int) int {
	return sl.buf[idx % len(sl.buf)]
}

func (sl *spinlock) position() int {
	return sl.pos
}

func (sl *spinlock) find(val int) int {
	for i, v := range sl.buf {
		if v == val {
			return i
		}
	}
	return -1
}

func (sl *spinlock) String() string {
	buf := sl.buf
	pos := sl.pos

	var b bytes.Buffer
	appendf := func(format string, a ...interface{}) {
		fmt.Fprintf(&b, format, a...)
	}

	v := buf[0]
	if pos == 0 {
		appendf("(%d)", v)
		
	} else {
		appendf("%d", v)
	}

	len := len(buf)
	for i := 1; i < len; i++ {
		v = buf[i]
		switch i {
			case pos: appendf(" (%d)", v)
			case pos+1 : appendf(" %d", v)
			default: appendf("  %d", v)
		}
	}

	return b.String()
}


// Day17 is the 17th day in Advent of Code.
func Day17() *day.Day {

	solve := func() (interface{}, interface{}) {
		sl := newSpinlock(343, 50000000)
		sl.advanceTo(2018)

		pt1 := sl.get(sl.position() + 1)

		//sl.advanceTo(50000000)
		
		//pt2 := sl.get(sl.find(0) + 1)
		
		return pt1, -1
	}

	return day.New(17, "Spinlock", solve)
}
