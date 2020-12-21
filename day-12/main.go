package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	inp := getInputs()
	start := time.Now()
	part1(inp)
	fmt.Println(time.Since(start))
	start = time.Now()
	part2(inp)
	fmt.Println(time.Since(start))
}
func part1(inp []string) {
	facing := 0
	x, y := 0, 0

	for _, line := range inp {
		dir := line[0]
		n, _ := strconv.Atoi(line[1:])

		if dir == 'N' {
			y += n
		} else if dir == 'W' {
			x -= n
		} else if dir == 'S' {
			y -= n
		} else if dir == 'E' {
			x += n
		} else if dir == 'F' {
			if facing == 0 {
				x += n
			} else if facing == 1 {
				y -= n
			} else if facing == 2 {
				x -= n
			} else if facing == 3 {
				y += n
			}
		} else if dir == 'L' {
			facing = mod(facing-n/90, 4)
		} else if dir == 'R' {
			facing = mod(facing+n/90, 4)
		}
	}
	fmt.Println("Part 1:", absint(x)+absint(y))
}
func part2(inp []string) {
	x, y := 0, 0
	relX, relY := 10, 1

	for _, line := range inp {
		dir := line[0]
		n, _ := strconv.Atoi(line[1:])

		if dir == 'N' {
			relY += n
		} else if dir == 'W' {
			relX -= n
		} else if dir == 'S' {
			relY -= n
		} else if dir == 'E' {
			relX += n
		} else if dir == 'F' {
			y += relY * n
			x += relX * n
		} else if dir == 'L' {
			for n > 0 {
				relX, relY = -relY, relX
				n -= 90
			}
		} else if dir == 'R' {
			for n > 0 {
				relX, relY = relY, -relX
				n -= 90
			}
		}
	}
	fmt.Println("Part 2:", absint(x)+absint(y))
}
func absint(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}
func getInputs() []string {
	data, _ := ioutil.ReadFile("./bigboy.txt")
	return strings.Split(string(data), "\n")
}
