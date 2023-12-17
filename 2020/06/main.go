package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {

	m := getQuestions()

	fmt.Println("Part 1:")
	part1(m)
	fmt.Println("\nPart 2:")
	part2(m)
}

func part1(m [][]string) {
	var lol []map[string]bool
	for _, l := range m {
		m := make(map[string]bool)
		for _, w := range l {
			for _, l := range w {
				m[string(l)] = true
			}
		}
		lol = append(lol, m)
	}

	sum := 0
	for _, v := range m {
		sum += len(v)
	}
	fmt.Println(sum)
}

func part2(m [][]string) {

	sum := 0
	for _, part := range m {
		sort.Slice(part, func(i, j int) bool { return len(part[i]) < len(part[j]) })
		start := part[0]
		if len(part) == 1 {
			sum += len(start)
			continue
		}
		ok := 0
		for _, l := range start {
			good := true
			for _, word := range part[1:] {
				if !strings.ContainsRune(word, l) {
					good = false
					break
				}
			}
			if good {
				ok++
			}
		}
		sum += ok
	}
	fmt.Println(sum)
}

func getQuestions() [][]string {
	data, _ := ioutil.ReadFile("./input.txt")
	var l [][]string
	for _, w := range strings.Split(string(data), "\n\n") {
		l = append(l, strings.Fields(w))
	}
	return l
}
