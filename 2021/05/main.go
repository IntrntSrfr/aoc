package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	size := 1000
	lines := getInputs()

	stats := &stats{m: map[int]int{}}

	now := time.Now()
	solve(lines, size, false, stats)
	fmt.Println("part 1:", stats.overlaps)
	fmt.Println(time.Since(now))

	now2 := time.Now()
	solve(lines, size, true, stats)
	fmt.Println("part 2:", stats.overlaps)
	fmt.Println(time.Since(now2))
	fmt.Println("total time:", time.Since(now))
}

type stats struct {
	m        map[int]int
	overlaps int
}

func solve(lines [][4]int, size int, part2 bool, s *stats) {
	for _, line := range lines {
		x0, y0, x1, y1 := line[0], line[1], line[2], line[3]
		if x0 == x1 && !part2 {
			for y := min(y0, y1); y < max(y0, y1)+1; y++ {
				if s.m[(size*y)+x0] == 1 {
					s.overlaps++
				}
				s.m[(size*y)+x0]++
			}
		} else if y0 == y1 && !part2 {
			for x := min(x0, x1); x < max(x0, x1)+1; x++ {
				if s.m[(size*y0)+x] == 1 {
					s.overlaps++
				}
				s.m[(size*y0)+x]++
			}
		} else if (abs(x0-x1) == abs(y0-y1)) && part2 {
			dx, dy := 1, 1
			if x1 < x0 {
				dx = -1
			}
			if y1 < y0 {
				dy = -1
			}
			for i := 0; i < abs(x0-x1)+1; i++ {
				ypos, xpos := y0+i*dy, x0+i*dx
				if s.m[size*(ypos)+(xpos)] == 1 {
					s.overlaps++
				}
				s.m[size*(ypos)+(xpos)]++
			}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getInputs() [][4]int {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines [][4]int
	for scanner.Scan() {
		var intCoords [4]int
		points := strings.ReplaceAll(scanner.Text(), " -> ", ",")
		coords := strings.Split(points, ",")
		x0, _ := strconv.Atoi(coords[0])
		y0, _ := strconv.Atoi(coords[1])
		x1, _ := strconv.Atoi(coords[2])
		y1, _ := strconv.Atoi(coords[3])
		intCoords[0] = x0
		intCoords[1] = y0
		intCoords[2] = x1
		intCoords[3] = y1
		lines = append(lines, intCoords)
	}
	return lines
}
