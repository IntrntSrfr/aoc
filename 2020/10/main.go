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

	jolts := map[int]int{1: 0, 3: 0}
	sort.Ints(inp)
	inp = append(inp, inp[len(inp)-1]+3)
	for i := 0; i < len(inp)-1; i++ {
		diff := inp[i+1] - inp[i]
		jolts[diff]++
	}

	fmt.Println(jolts[1] * jolts[3])

	c := map[int]int{0: 1}
	for _, v := range inp {
		c[v] += c[v-1] + c[v-2] + c[v-3]
	}

	fmt.Println(c[inp[len(inp)-1]])
}

func getInputs() []int {
	data, _ := ioutil.ReadFile("./input.txt")
	res := []int{0}
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		res = append(res, n)
	}
	return res
}
