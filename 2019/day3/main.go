package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	m := make(map[string]bool)
	dir := getDirections()

	x, y := 0, 0
	for _, v := range dir[0] {

		direction := string(v[0])
		lenght, _ := strconv.Atoi(v[1:])

		switch direction {
		case "R":
			for i := 0; i < lenght; i++ {
				x1, y1 := strconv.Itoa(x+i), strconv.Itoa(y)
				m[x1+","+y1] = true
			}
			x += lenght
		case "L":
			for i := 0; i < lenght; i++ {
				x1, y1 := strconv.Itoa(x-i), strconv.Itoa(y)
				m[x1+","+y1] = true
			}
			x -= lenght
		case "U":
			for i := 0; i < lenght; i++ {
				x1, y1 := strconv.Itoa(x), strconv.Itoa(y+i)
				m[x1+","+y1] = true
			}
			y += lenght
		case "D":
			for i := 0; i < lenght; i++ {
				x1, y1 := strconv.Itoa(x), strconv.Itoa(y-i)
				m[x1+","+y1] = true
			}
			y -= lenght
		}
	}

	intersections := make(map[string]bool)

	x, y = 0, 0
	for _, v := range dir[1] {

		direction := string(v[0])
		lenght, _ := strconv.Atoi(v[1:])

		switch direction {
		case "R":
			for i := 0; i < lenght; i++ {
				x1, y1 := strconv.Itoa(x+i), strconv.Itoa(y)
				if _, ok := m[x1+","+y1]; ok {
					intersections[x1+","+y1] = true
				}
				m[x1+y1] = true
			}
			x += lenght
		case "L":
			for i := 0; i < lenght; i++ {
				x1, y1 := strconv.Itoa(x-i), strconv.Itoa(y)
				if _, ok := m[x1+","+y1]; ok {
					intersections[x1+","+y1] = true
				}
				m[x1+y1] = true
			}
			x -= lenght
		case "U":
			for i := 0; i < lenght; i++ {
				x1, y1 := strconv.Itoa(x), strconv.Itoa(y+i)
				if _, ok := m[x1+","+y1]; ok {
					intersections[x1+","+y1] = true
				}
				m[x1+y1] = true
			}
			y += lenght
		case "D":
			for i := 0; i < lenght; i++ {
				x1, y1 := strconv.Itoa(x), strconv.Itoa(y-i)
				if _, ok := m[x1+","+y1]; ok {
					intersections[x1+","+y1] = true
				}
				m[x1+y1] = true
			}
			y -= lenght
		}
	}

	lowest := 10000000000.0
	for i := range intersections {
		s := strings.Split(i, ",")
		x, _ := strconv.ParseFloat(s[0], 64)
		y, _ := strconv.ParseFloat(s[1], 64)

		l := math.Abs(0-x) + math.Abs(0-y)
		if l < lowest && l != 0 {
			lowest = l
		}
	}
	fmt.Println(lowest)
}

func getDirections() [][]string {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	instructions := [][]string{}
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		i := strings.Split(scan.Text(), ",")

		instructions = append(instructions, i)
	}

	return instructions
}

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}
