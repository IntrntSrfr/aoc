package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {

	grid := getGrid()

	res := []int{0, 0, 0}
	resmap := make(map[[2]int]bool)

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "." {
				continue
			}

			seen := make(map[[2]int]bool)

			for y2 := range grid {
				for x2 := range grid[y2] {
					if grid[y2][x2] == "#" {
						if y == y2 && x == x2 {
							continue
						}
						dx, dy := x2-x, y2-y

						g := int(math.Abs(float64(gcd(dx, dy))))
						dx /= g
						dy /= g

						if _, ok := seen[[2]int{dy, dx}]; !ok {
							seen[[2]int{dy, dx}] = true
						}
					}
				}
			}

			if len(seen) > res[0] {
				res = []int{len(seen), y, x}
				resmap = seen
			}
		}
	}
	fmt.Println("part 1:", res)

	//fmt.Println(resmap)

	t := []tar{}

	for i := range resmap {
		key := 180 - ((180 / math.Pi) * math.Atan2(float64(i[1]), float64(i[0])))

		t = append(t, tar{angle: key, y: res[1] + i[0], x: res[2] + i[1], dy: i[0], dx: i[1]})
	}

	sort.Sort(sort.Reverse(SortByAngle(t)))

	//fmt.Println(t)
	target := t[199]
	fmt.Println(target)

}

func gcd(x, y int) int {
	for y != 0 {
		//fmt.Println(x, y)
		x, y = y, x%y
	}
	return x
}

func getGrid() [][]string {

	f, _ := os.Open("./input.txt")
	defer f.Close()

	res := [][]string{}
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		row := []string{}
		for _, v := range scan.Text() {
			row = append(row, string(v))
		}
		res = append(res, row)
	}

	return res
}

type tar struct {
	angle float64
	y     int
	x     int
	dy    int
	dx    int
}

type SortByAngle []tar

func (a SortByAngle) Len() int           { return len(a) }
func (a SortByAngle) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAngle) Less(i, j int) bool { return a[i].angle < a[j].angle }
