package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp := getInputs()
	part1(inp)
	part2(inp)
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
	fmt.Println(count)
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
			bitset := (four & (0b1111111 ^ seven))
			if bits.OnesCount(bitset&n) == 1 {
				b = (bitset & n)
				d = (bitset & (0b1111111 ^ n))
			}
			g &= n
		}
		transfer := g
		g &= (0b1111111 ^ (four | a))
		e := eight & (0b1111111 ^ (four | a | g))
		c := eight & (0b1111111 ^ (transfer | e | d))
		f := eight & (0b1111111 ^ (a | b | c | d | e | g))

		mapping := map[uint]string{
			a | b | c | e | f | g:     "0",
			c | f:                     "1",
			a | c | d | e | g:         "2",
			a | c | d | f | g:         "3",
			b | c | d | f:             "4",
			a | b | d | f | g:         "5",
			a | b | d | e | f | g:     "6",
			a | c | f:                 "7",
			a | b | c | d | e | f | g: "8",
			a | b | c | d | f | g:     "9",
		}

		str := ""
		for _, ye := range line[10:] {
			str += mapping[ye]
		}
		num, _ := strconv.Atoi(str)
		sum += num
	}
	fmt.Println(sum)
}

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
			res += (1 << 6)
		case 'b':
			res += (1 << 5)
		case 'c':
			res += (1 << 4)
		case 'd':
			res += (1 << 3)
		case 'e':
			res += (1 << 2)
		case 'f':
			res += (1 << 1)
		case 'g':
			res += 1
		}
	}
	return res
}

func getInputs() [][]uint {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

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
