package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
	"time"
)

func main() {
	inp := getInputs()
	p1Start := time.Now()
	part1(inp)
	fmt.Println("part 1 time:  ", time.Since(p1Start))
	p2Start := time.Now()
	part2(inp)
	fmt.Println("part 2 time:  ", time.Since(p2Start))
	fmt.Println("total time:   ", time.Since(p1Start))
}

func part1(inp [][]uint) {
	count := 0
	for _, line := range inp {
		for _, code := range line[10:] {
			n := bits.OnesCount(code)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				count++
			}
		}
	}
	fmt.Println("part 1 answer:", count)
}

func part2(inp [][]uint) {
	sum := 0
	for _, line := range inp {
		one, seven, four, eight := cwbc(line[:10], 2)[0], cwbc(line[:10], 3)[0], cwbc(line[:10], 4)[0], cwbc(line[:10], 7)[0]
		a := (one | seven) & (0b1111111 ^ one)
		var b uint
		var d uint
		var g uint = 0b1111111
		for _, n := range cwbc(line[:10], 6) {
			bitset := four & (0b1111111 ^ seven)
			if bits.OnesCount(bitset&n) == 1 {
				b = bitset & n
				d = bitset & (0b1111111 ^ n)
			}
			g &= n
		}
		transfer := g
		g &= 0b1111111 ^ (four | a)
		e := eight & (0b1111111 ^ (four | a | g))
		c := eight & (0b1111111 ^ (transfer | e | d))
		f := eight & (0b1111111 ^ (a | b | c | d | e | g))

		mapping := map[uint]int{
			a | b | c | e | f | g:     0,
			c | f:                     1,
			a | c | d | e | g:         2,
			a | c | d | f | g:         3,
			b | c | d | f:             4,
			a | b | d | f | g:         5,
			a | b | d | e | f | g:     6,
			a | c | f:                 7,
			a | b | c | d | e | f | g: 8,
			a | b | c | d | f | g:     9,
		}

		base := 1000
		num := 0
		for _, ye := range line[10:] {
			num += mapping[ye] * base
			base /= 10
		}
		sum += num
	}
	fmt.Println("part 2 answer:", sum)
}

// finds codes in list that has `count` amount of ones
func cwbc(a []uint, count int) []uint {
	var res []uint
	for _, i := range a {
		if bits.OnesCount(i) == count {
			res = append(res, i)
		}
	}
	return res
}

func codeToInt(code string) uint {
	var res uint
	for _, c := range code {
		switch c {
		case 'a':
			res += 1 << 6
		case 'b':
			res += 1 << 5
		case 'c':
			res += 1 << 4
		case 'd':
			res += 1 << 3
		case 'e':
			res += 1 << 2
		case 'f':
			res += 1 << 1
		case 'g':
			res += 1
		}
	}
	return res
}

func getInputs() [][]uint {
	scanner := bufio.NewScanner(os.Stdin)
	var lines [][]uint
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "|", "")
		fields := strings.Fields(line)
		var ints []uint
		for _, numStr := range fields {
			ints = append(ints, codeToInt(numStr))
		}
		lines = append(lines, ints)
	}
	return lines
}
