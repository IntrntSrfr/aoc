package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	m := make(map[string]string)

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		i := strings.Split(scan.Text(), ")")
		base, orbit := i[0], i[1]
		m[orbit] = base
	}

	c := 0
	for i := range m {
		c += traverse(m, i)
	}

	fmt.Println("part 1:", c-len(m))
	fmt.Println("part 2:", travel(m, "YOU", "SAN"))
}

func travel(m map[string]string, start, finish string) int {
	b := make(map[string]bool)

	c := 0
	node := start
	for {
		ok := true

		n1 := node
		node, ok = m[node]
		if !ok {
			break
		}
		c++
		n2 := node
		b[n1+n2] = true
	}

	node = finish
	for {
		ok := true

		n1 := node
		node, ok = m[node]
		if !ok {
			break
		}

		n2 := node
		if _, ok := b[n1+n2]; ok {
			c--
		} else {
			c++
		}
	}

	return c - 2

}

func traverse(m map[string]string, start string) int {

	node := start
	c := 0
	for {
		c++
		ok := true
		node, ok = m[node]
		if !ok {
			break
		}
	}
	return c
}
