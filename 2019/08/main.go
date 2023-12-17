package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func main() {
	inp := getList()

	w, h := 25, 6
	layers := [100][150]int{}

	for y := range layers {
		for x := range layers[y] {
			pos := x + (y * (w * h))
			layers[y][x] = inp[pos]
		}
	}

	layer := [150]int{}

	minZeros := math.MaxInt32
	for _, v := range layers {
		zeros := 0
		for j := range v {
			if v[j] == 0 {
				zeros++
			}
		}
		if zeros < minZeros {
			minZeros = zeros
			layer = v
		}
	}

	m := make(map[int]int)
	for _, v := range layer {
		m[v]++
	}

	fmt.Println("part 1:", m[1]*m[2])

	fmt.Println("part 2:")
	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			for l := 0; l < len(inp)/(w*h); l++ {
				if r := inp[l*w*h+y*w+x]; r != 2 {
					fmt.Print(map[int]string{0: " ", 1: "#"}[r])
					break
				}
			}
		}
		fmt.Println("")
	}
}

func getList() []int {

	f, _ := os.Open("./input.txt")
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	v := string(b)

	v2 := []int{}
	for i := range v {
		n, _ := strconv.Atoi(string(v[i]))
		v2 = append(v2, n)
	}

	return v2
}
