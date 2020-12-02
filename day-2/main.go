package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pws := getPasswords()

	fmt.Println("Part 1:")
	part1(pws)
	fmt.Println("Part 2:")
	part2(pws)
}

func part1(inp []Password) {

	good := 0
	for _, p := range inp {
		count := 0
		for _, l := range p.password {
			if l == p.letter {
				count++
			}
		}
		if count < p.min || p.max < count {
			continue
		}
		good++
	}

	fmt.Println("Valid passwords:", good)
}

func part2(inp []Password) {

	good := 0
	for _, p := range inp {
		if p.letter == rune(p.password[p.min-1]) && p.letter != rune(p.password[p.max-1]) {
			good++
		} else if p.letter == rune(p.password[p.max-1]) && p.letter != rune(p.password[p.min-1]) {
			good++
		}
	}

	fmt.Println("Vald passwords:", good)
}

func getPasswords() []Password {
	f, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(f)
	var pws []Password
	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, " ")

		lengths := strings.Split(parts[0], "-")
		min, _ := strconv.Atoi(lengths[0])
		max, _ := strconv.Atoi(lengths[1])

		pw := Password{
			min:      min,
			max:      max,
			letter:   rune(parts[1][0]),
			password: parts[2],
		}

		pws = append(pws, pw)
	}
	return pws
}

type Password struct {
	min      int
	max      int
	letter   rune
	password string
}
