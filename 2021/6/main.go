package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	state := getInputs()
	solve(state, 80)
	solve(state, 256)
}

func solve(state map[int]int, days int) {
	for day := 0; day < days; day++ {
		newState := make(map[int]int)
		for i, n := range state {
			if i == 0 {
				newState[6] += n
				newState[8] += n
				continue
			}
			newState[i-1] += n
		}
		state = newState
	}
	fmt.Println(sum(state))
}

func getInputs() map[int]int {
	f, _ := os.Open("./input.txt")
	line, _ := io.ReadAll(f)
	res := make(map[int]int)
	for _, numStr := range strings.Split(string(line), ",") {
		num, _ := strconv.Atoi(numStr)
		res[num]++
	}
	return res
}

func sum(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}
