package main

import (
	"bufio"
	"fmt"
	"github.com/intrntsrfr/aoc/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	grid, folds := getInputs()
	start := time.Now()
	for i, fold := range folds {
		grid = foldGrid(grid, fold)
		if i == 0 {
			fmt.Println("part 1:", utils.Sum2DGrid(grid))
		}
	}
	taken := time.Since(start)
	fmt.Println(len(grid), len(grid[0]))
	display(grid)
	fmt.Println("time taken:", taken)

}

func foldGrid(grid [][]int, axis string) [][]int {
	var newGrid [][]int
	if axis == "y" {
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

func getInputs() ([][]int, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	maxY, _ := strconv.Atoi(os.Args[1])
	maxX, _ := strconv.Atoi(os.Args[2])

	grid := make([][]int, maxY)
	for y := range grid {
		grid[y] = make([]int, maxX)
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

	var folds []string
	for scanner.Scan() {
		ye := strings.Split(scanner.Text(), "=")
		folds = append(folds, string(ye[0][len(ye[0])-1]))
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
