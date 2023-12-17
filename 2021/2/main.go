package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution(true)
	solution(false)
}

func solution(part1 bool) {
	aim, x, y := 0, 0, 0

	inp := getInputs()

	for _, v := range inp {
		if v[0] == 0 {
			// if forward
			x += v[1]
			if !part1 {
				y += (aim * v[1])
			}
		} else {
			// if down/up
			if !part1 {
				aim += v[1]
			} else {
				y += v[1]
			}
		}
	}

	t := "part 1:"
	if !part1 {
		t = "part 2:"
	}
	fmt.Println(t, x*y)
}

func getInputs() [][]int {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var res [][]int
	for scanner.Scan() {
		ye := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(ye[1])

		line := make([]int, 2)

		if ye[0] == "forward" {
			line[0], line[1] = 0, n
		} else if ye[0] == "up" {
			line[0], line[1] = 1, -n
		} else {
			line[0], line[1] = 1, n
		}

		res = append(res, line)
	}

	return res
}
