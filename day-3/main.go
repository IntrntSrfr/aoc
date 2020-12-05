package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inp := getInput()

	fmt.Println("Part 1:")
	part1(inp)
	fmt.Println("Part 2:")
	part2(inp)
}

func part1(inp Trees) {
	trees := findTrees(inp, 1, 3)
	fmt.Println("Trees found:", trees)
}

func part2(inp Trees) {
	moves := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	sum := 1
	for _, move := range moves {
		sum *= findTrees(inp, move[0], move[1])
	}
	fmt.Println("Multiplied trees encountered:", sum)
}
func findTrees(inp Trees, dx, dy int) int {
	found := 0
	for x, y := 0, 0; x < len(inp.trees); x, y = x+dx, y+dy {
		if inp.trees[x][y%len(inp.trees[x])] == 1 {
			found++
		}
	}
	return found
}

func getInput() Trees {

	f, _ := os.Open("./input.txt")
	scan := bufio.NewScanner(f)

	t := Trees{
		width: 0,
		trees: [][]int{},
	}

	for scan.Scan() {
		line := scan.Text()
		t.width = len(line)

		var treeline []int
		for _, l := range line {
			if l == '#' {
				treeline = append(treeline, 1)
			} else {
				treeline = append(treeline, 0)
			}
		}
		t.trees = append(t.trees, treeline)
	}
	return t
}

type Trees struct {
	width int
	trees [][]int
}
