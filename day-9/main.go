package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inp := getInputs()

	ans1 := part1(inp, 25)
	fmt.Println(ans1)

	ans2 := part2(inp, ans1)
	fmt.Println(ans2)
}

func part1(inp []int, preamble int) int {
	for i := preamble; i < len(inp)-1; i++ {
		any := false
	combination:
		for m := i - preamble; m < i; m++ {
			for n := m + 1; n < i; n++ {
				if inp[m]+inp[n] == inp[i] {
					any = true
					break combination
				}
			}
		}
		if !any {
			return inp[i]
		}
	}
	return -1
}

func part2(inp []int, target int) int {
	set := make([]int, 0)
setLoop:
	for m := 0; m < len(inp); m++ {
		set = []int{}
		sum := 0
		for n := m; n < len(inp); n++ {
			sum += inp[n]
			set = append(set, inp[n])
			if sum == target {
				break setLoop
			}
		}
	}
	sort.Ints(set)
	return set[0] + set[len(set)-1]
}

func getInputs() []int {
	data, _ := ioutil.ReadFile("./input.txt")
	var res []int
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		res = append(res, n)
	}
	return res
}
