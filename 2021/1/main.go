package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	inp := getInputs()
	part1(inp)
	part2(inp)
}

func part1(inp []int) {
	inc := 0
	for i := 1; i < len(inp); i++ {
		if inp[i-1] < inp[i] {
			inc++
		}
	}

	fmt.Println("part 1:", inc)
}

func part2(inp []int) {
	inc := 0
	lastSum := inp[0] + inp[1] + inp[2]
	for i := 1; i < len(inp)-1; i++ {
		sum := inp[i-1] + inp[i] + inp[i+1]
		if sum > lastSum {
			inc++
		}
		lastSum = sum
	}

	fmt.Println("part 2:", inc)
}

func getInputs() []int {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var res []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		res = append(res, n)
	}

	return res
}
