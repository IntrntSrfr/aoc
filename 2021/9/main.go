package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	grid := getInputs()

	p1Start := time.Now()
	minima := findMinima(grid)
	fmt.Println("part 1 time:", time.Since(p1Start))

	p2Start := time.Now()
	part2(grid, minima)
	fmt.Println("part 2 time:", time.Since(p2Start))
	fmt.Println("total time:", time.Since(p1Start))
}

func findMinima(grid [][]int) [][]int {
	minima := [][]int{}
	for y := range grid {
		for x := range grid[y] {
			v := grid[y][x]
			if y > 0 && v >= grid[y-1][x] {
				continue
			}
			if y < len(grid)-1 && v >= grid[y+1][x] {
				continue
			}
			if x > 0 && v >= grid[y][x-1] {
				continue
			}
			if x < len(grid[y])-1 && v >= grid[y][x+1] {
				continue
			}
			minima = append(minima, []int{y, x, v})
		}
	}
	fmt.Println("part 1 answer:", sumCoords(minima)+len(minima))
	return minima
}

func part2(grid [][]int, minima [][]int) {
	width := len(grid[0])
	var basins []int
	for _, minimum := range minima {
		visited := make(map[int]bool)
		pos := minimum[0]*width + minimum[1]
		visited[pos] = true
		size := 0
		q := NewQueue()
		q.Push(minimum)
		for !q.IsEmpty() {
			ye := q.Pop()
			size++
			y, x := ye[0], ye[1]

			// edge cases can be ignored(?) if padding is added

			if y > 0 && grid[y-1][x] != 9 {
				if !visited[(y-1)*width+x] {
					visited[(y-1)*width+x] = true
					q.Push([]int{y - 1, x})
				}
			}
			if y < len(grid)-1 && grid[y+1][x] != 9 {
				if !visited[(y+1)*width+x] {
					visited[(y+1)*width+x] = true
					q.Push([]int{y + 1, x})
				}
			}
			if x > 0 && grid[y][x-1] != 9 {
				if !visited[y*width+x-1] {
					visited[y*width+x-1] = true
					q.Push([]int{y, x - 1})
				}
			}
			if x < len(grid[y])-1 && grid[y][x+1] != 9 {
				if !visited[y*width+x+1] {
					visited[y*width+x+1] = true
					q.Push([]int{y, x + 1})
				}
			}
		}
		basins = append(basins, size)
	}
	sort.Ints(basins)
	fmt.Println("part 2 answer:", product(basins[len(basins)-3:]))
}

func product(a []int) int {
	sum := 1
	for _, v := range a {
		sum *= v
	}
	return sum
}

func sumCoords(g [][]int) int {
	sum := 0
	for _, v := range g {
		sum += v[2]
	}
	return sum
}

func getInputs() [][]int {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var ints []int
		for _, c := range line {
			ints = append(ints, int(c-'0'))
		}
		grid = append(grid, ints)
	}
	return grid
}

type Queue struct {
	data [][]int
}

func NewQueue() *Queue {
	return &Queue{
		data: [][]int{},
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Push(v []int) {
	q.data = append(q.data, v)
}

func (q *Queue) Pop() []int {
	if q.IsEmpty() {
		return []int{}
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
