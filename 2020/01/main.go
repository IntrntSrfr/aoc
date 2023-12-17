package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	inp := getInputs()

	sort.Ints(inp)

	fmt.Println("part 1:")
	part1(inp)
	fmt.Println("part 2:")
	part2(inp)
}

func part1(inp []int) {
	start := time.Now()
loop:
	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
			if inp[i]+inp[j] == 2020 {
				fmt.Println(inp[i], inp[j], inp[i]*inp[j])
				break loop
			}
		}
	}
	fmt.Println("time taken:", time.Now().Sub(start))
}

func part2(inp []int) {
	start := time.Now()

loop:
	for i := 0; i < len(inp); i++ {
		for j := i + 1; j < len(inp); j++ {
			for k := j + 1; k < len(inp); k++ {
				if inp[i]+inp[j]+inp[k] == 2020 {
					fmt.Println(inp[i], inp[j], inp[k], inp[i]*inp[j]*inp[k])
					break loop
				}
			}
		}
	}

	fmt.Println("time taken:", time.Now().Sub(start).String())
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
