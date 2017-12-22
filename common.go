package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	check(err)
	return i
}

func readInput(day int) []byte {
	path := fmt.Sprintf("./data/day%d/input.txt", day)
	dat, err := ioutil.ReadFile(path)
	check(err)
	return dat
}

func openInput(day int) *os.File {
	path := fmt.Sprintf("./data/day%d/input.txt", day)
	f, err := os.Open(path)
	check(err)
	return f
}

func readInputLine(day int) string {
	file := openInput(day)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		return line
	}

	panic("Uh-oh!")
}
