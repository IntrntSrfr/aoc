package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inp := getInput()
	//fmt.Println(inp)

	fmt.Println("Part 1:")
	part1(inp)
	fmt.Println("Part 2:")
	part2(inp)
}

func part1(inp Trees) {
	trees := 0
	for x, y := 0, 0; x < len(inp.trees); x, y = x+1, y+3 {
		if inp.trees[x][y%len(inp.trees[x])] == 1 {
			trees++
		}
	}
	fmt.Println("Trees found:", trees)
}

func part2(inp Trees) {

	moves := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	var res []int

	for _, move := range moves {
		found := 0
		for x, y := 0, 0; x < len(inp.trees); x, y = x+move[0], y+move[1] {
			if inp.trees[x][y%len(inp.trees[x])] == 1 {
				found++
			}
		}
		res = append(res, found)
	}

	sum := 1
	for i := 0; i < len(res); i++ {
		sum *= res[i]
	}

	fmt.Println("Multiplied trees encountered:", sum)
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
