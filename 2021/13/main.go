package main

import (
	"bufio"
	"fmt"
	"github.com/intrntsrfr/aoc/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid, folds := getInputs()
	for i, fold := range folds {
		grid = foldGrid(grid, fold[0], fold[1])
		if i == 0 {
			fmt.Println("part 1:", utils.Sum2DGrid(grid))
		}
	}
	display(grid)
}

func foldGrid(grid [][]int, foldY int, value int) [][]int {
	var newGrid [][]int
	if foldY == 1 {
		newGrid = utils.MxNGrid(len(grid)/2, len(grid[0])) // fold y axis
		for y := range newGrid {
			for x := range newGrid[y] {
				newGrid[y][x] = grid[y][x] | grid[len(grid)-y-1][x]
			}
		}
	} else {
		newGrid = utils.MxNGrid(len(grid), len(grid[0])/2) // fold x axis
		for y := range newGrid {
			for x := range newGrid[y] {
				newGrid[y][x] = grid[y][x] | grid[y][len(grid[y])-x-1]
			}
		}
	}
	return newGrid
}

func getInputs() ([][]int, [][]int) {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]int, 55)
	for y := range grid {
		grid[y] = make([]int, 30)
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		coords := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x] = 1
	}

	var folds [][]int
	for scanner.Scan() {
		ye := strings.Split(scanner.Text(), "=")
		value, _ := strconv.Atoi(ye[1])
		if ye[0][len(ye[0])-1] == 'x' {
			folds = append(folds, []int{0, value})
			continue
		}
		folds = append(folds, []int{1, value})
	}
	return grid, folds
}

func display(grid [][]int) {
	for y := range grid {
		row := ""
		for x := range grid[y] {
			if grid[y][x] == 1 {
				row += "██"
			} else {
				row += "  "
			}
		}
		fmt.Println(row)
	}
}
