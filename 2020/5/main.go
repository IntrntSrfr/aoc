package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	var IDs []int
	f, _ := os.Open("./input.txt")
	defer f.Close()
	r := bufio.NewReader(f)

	p1 := time.Now()
	highest := 0
	for {
		pass, _, err := r.ReadLine()
		if err != nil {
			break
		}
		n := 0
		for _, b := range pass {
			if b == byte('F') || b == byte('L') {
				n = n << 1
			} else {
				n = (n << 1) + 1
			}
		}
		if n > highest {
			highest = n
		}
		IDs = append(IDs, n)
	}
	fmt.Println("Part 1:", highest)
	fmt.Println("Part 1 time:", time.Since(p1))
	fmt.Println()

	p2 := time.Now()
	sort.Ints(IDs)
	for i := 0; i < len(IDs)-1; i++ {
		if IDs[i+1] != IDs[i]+1 {
			fmt.Println("Part 2:", IDs[i]+1)
			break
		}
	}
	fmt.Println("Part 2 time:", time.Since(p2))

	fmt.Println()
	fmt.Println("Full time:", time.Since(start))
	return
}
