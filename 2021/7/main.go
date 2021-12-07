package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp := getInputs()
	solve(inp, true)
	solve(inp, false)
}

func solve(inp []int, part1 bool) {
	minSum := 0
	for i := 0; i < 10000; i++ {
		sum := 0
		for _, x := range inp {
			if part1 {
				sum += abs(x - i)
			} else {
				sum += (abs(x-i) * (abs(x-i) + 1)) / 2
			}
		}
		if sum < minSum || minSum == 0 {
			minSum = sum
		}
	}
	fmt.Println(minSum)
}

func getInputs() []int {
	f, _ := os.Open("./input.txt")
	line, _ := io.ReadAll(f)
	var res []int
	for _, numStr := range strings.Split(string(line), ",") {
		num, _ := strconv.Atoi(numStr)
		res = append(res, num)
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
