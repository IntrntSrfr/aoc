package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	inp := getInputs()

	start := time.Now()
	runSeats(inp, false)
	fmt.Println(time.Since(start))
	start = time.Now()
	runSeats(inp, true)
	fmt.Println(time.Since(start))
}

func runSeats(inp [][]string, p2 bool) {
	for {
		count := 0
		newSeats := copySeats(inp)
		diff := false

		for y := 0; y < len(inp); y++ {
			for x := 0; x < len(inp[y]); x++ {
				if inp[y][x] == "." {
					continue
				}

				n := cn(inp, y, x, p2)

				if inp[y][x] == "L" && n == 0 {
					newSeats[y][x] = "#"
					diff = true
				} else if inp[y][x] == "#" {
					if p2 && n >= 5 {
						newSeats[y][x] = "L"
						diff = true
					} else if !p2 && n >= 4 {
						newSeats[y][x] = "L"
						diff = true
					}
				}
				if newSeats[y][x] == "#" {
					count++
				}
			}
		}

		if !diff {
			if !p2 {
				fmt.Println("Part 1:")
				fmt.Println(count)
			} else {
				fmt.Println("Part 2:")
				fmt.Println(count)
			}
			break
		}
		inp = newSeats
	}
}

func cn(l [][]string, y, x int, long bool) int {
	neighbours := 0
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			length := 1
			for {
				if y+(dy*length) < 0 || x+(dx*length) < 0 || y+(dy*length) > len(l)-1 || x+(dx*length) > len(l[y])-1 {
					break
				}
				if l[y+(dy*length)][x+(dx*length)] == "#" {
					neighbours++
					break
				}
				if long {
					if l[y+(dy*length)][x+(dx*length)] == "L" {
						break
					}
					length++
				} else {
					break
				}
			}
		}
	}
	return neighbours
}

func copySeats(l [][]string) [][]string {
	res := make([][]string, len(l))
	for i, r := range l {
		ll := make([]string, len(l[i]))
		for j, c := range r {
			ll[j] = c
		}
		res[i] = ll
	}
	return res
}

func getInputs() [][]string {
	data, _ := ioutil.ReadFile("./bigboy.txt")
	var res [][]string
	for _, line := range strings.Split(string(data), "\n") {
		var l []string
		for _, c := range line {
			l = append(l, string(c))
		}
		res = append(res, l)
	}
	return res
}
