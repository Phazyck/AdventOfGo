package main

import (
	"bufio"
	"regexp"
	"strings"
)

type tower struct {
	below  map[string]string
	above  map[string][]string
	weight map[string]int
	root   string
}

func (t *tower) add(name string, weight int, above []string) {
	t.weight[name] = weight

	if above != nil {
		t.above[name] = above
		for _, n := range above {
			t.below[n] = name
		}
	}
}

func (t *tower) find(name string) (weight int, found bool) {
	names, ok := t.above[name]
	weight = t.weight[name]
	if !ok {
		return weight, false
	}

	weights := make(map[int][]string)

	for _, n := range names {
		w, found := t.find(n)
		if found {
			return w, true
		}

		weight += w

		if v, ok := weights[w]; !ok {
			weights[w] = []string{n}
		} else {
			weights[w] = append(v, n)
		}
	}

	if len(names) < 3 {
		return weight, false
	}

	var valMin, valMax []string
	var keyMin, keyMax int

	for key, val := range weights {
		if key > keyMax || valMax == nil {
			keyMax, valMax = key, val
		}

		if key < keyMin || valMin == nil {
			keyMin, valMin = key, val
		}
	}

	if keyMax != keyMin {
		var name string
		var actual, target int
		if len(valMax) == 1 {
			name = valMax[0]
			actual = keyMax
			target = keyMin
		} else {
			name = valMin[0]
			actual = keyMin
			target = keyMax
		}
		change := target - actual
		current := t.weight[name]
		return current + change, true
	}

	return weight, false
}

func day07() *day {

	build := func() *tower {
		f := openInput(7)

		scanner := bufio.NewScanner(f)

		tower := &tower{
			make(map[string]string),
			make(map[string][]string),
			make(map[string]int),
			"",
		}

		reNameAndWeight := regexp.MustCompile("(.*) \\((.*)\\)")
		reChildren := regexp.MustCompile("-> (.*)")

		candidates := make(map[string]bool)

		for scanner.Scan() {
			line := scanner.Text()

			mNameAndWeight := reNameAndWeight.FindStringSubmatch(line)
			name := mNameAndWeight[1]
			weight := parseInt(mNameAndWeight[2])

			candidates[name] = true

			mChildren := reChildren.FindStringSubmatch(line)
			var above []string
			if len(mChildren) == 2 {
				above = strings.Split(mChildren[1], ", ")
				for _, n := range above {
					delete(candidates, n)
				}
			}

			tower.add(name, weight, above)
		}

		for k := range tower.below {
			delete(candidates, k)
		}

		if len(candidates) != 1 {
			panic("This is unexptected...")
		}

		for k := range candidates {
			tower.root = k
		}

		return tower
	}

	solve := func() (interface{}, interface{}) {
		tower := build()
		part1 := tower.root
		part2, _ := tower.find(tower.root)
		return part1, part2
	}

	return &day{7, "Recursive Circus", solve}
}
