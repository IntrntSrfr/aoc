package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/intrntsrfr/aoc/utils"
)

/*
The energy level of each octopus is a value between 0 and 9. Here, the top-left octopus has an energy level of 5, the
bottom-right one has an energy level of 6, and so on.

You can model the energy levels and flashes of light in steps. During a single step, the following occurs:

- First, the energy level of each octopus increases by 1.
- Then, any octopus with an energy level greater than 9 flashes. This increases the energy level of all adjacent
  octopuses by 1, including octopuses that are diagonally adjacent. If this causes an octopus to have an energy
  level greater than 9, it also flashes. This process continues as long as new octopuses keep having their energy
  level increased beyond 9. (An octopus can only flash at most once per step.)
- Finally, any octopus that flashed during this step has its energy level set to 0, as it used all of its energy to flash.
Adjacent flashes can cause an octopus to flash on a step even if it begins that step with very little energy.

*/

func main() {
	grid := getInputs()

	//bufGrid := deepcopy(grid)
	utils.DisplayGrid(grid)

	for t := 0; t < 2; t++ {
		// increase every value by 1
		for y := range grid {
			for x := range grid[y] {
				grid[y][x]++
			}
		}

		visitMap := make(map[string]bool)

		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] > 9 {
					grid[y][x] = 0
					visitMap[makeKey(y, x)] = true
					increaseNeighbours(grid, y, x, visitMap)

					// increase neighbours by 1
					// if neighbour > 9 then do this step for them too
					// this can happen only once for each cell

				}
			}
		}

		utils.DisplayGrid(grid)
	}

	//fmt.Println(grid)

}

func makeKey(y, x int) string {
	return fmt.Sprintf("%v-%v", y, x)
}

func increaseNeighbours(grid [][]int, x, y int, visitMap map[string]bool) {
	// 0 = y, 1 = x
	moves := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, move := range moves {
		newY, newX := move[0]+y, move[1]+x
		// assume nxn grid
		if newY < 0 || newX < 0 || newY > len(grid)-1 || newX > len(grid) {
			continue
		}
		if !visitMap[makeKey(newY, newX)] {

		}
		grid[newY][newX]++
		if grid[newY][newX] > 9 {
			grid[newY][newX] = 0
			visitMap[makeKey(newY, newX)] = true
			increaseNeighbours(grid, newY, newX, visitMap)
		}
	}
}

func deepcopy(g [][]int) [][]int {
	var res [][]int
	for y := range g {
		row := make([]int, len(g[y]))
		copy(row, g[y])
		res = append(res, row)
	}
	return res
}

func getInputs() [][]int {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var res [][]int
	for scanner.Scan() {
		line := scanner.Text()

		var lineRes []int

		for _, c := range line {
			lineRes = append(lineRes, int(c)-48)
		}

		res = append(res, lineRes)

	}
	return res
}
