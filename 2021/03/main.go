package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inp := getInputs()
	part1(inp)
	part2(inp)
}

func part1(inp []string) {
	gr, lr := "", ""

	for i := 0; i < len(inp[0]); i++ {
		common := mostCommonAtPos(inp, i)
		if common == '0' {
			gr, lr = gr+"0", lr+"1"
		} else {
			gr, lr = gr+"1", lr+"0"
		}
	}

	gri, _ := strconv.ParseInt(gr, 2, 64)
	lri, _ := strconv.ParseInt(lr, 2, 64)

	fmt.Println("part 1:", gri*lri)
}

func part2(inp []string) {
	oxygenString := getThing(inp, true)
	co2String := getThing(inp, false)
	oxygen, _ := strconv.ParseInt(oxygenString, 2, 64)
	co2, _ := strconv.ParseInt(co2String, 2, 64)
	fmt.Println("part 2:", oxygen*co2)
}

func mostCommonAtPos(inp []string, pos int) byte {
	zeros, ones := 0, 0
	for _, line := range inp {
		if line[pos] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	if zeros > ones {
		return '0'
	}
	return '1'
}

func getThing(inp []string, most bool) string {
	for i := 0; i < len(inp[0]); i++ {
		if len(inp) == 1 {
			break
		}
		common := mostCommonAtPos(inp, i)

		if (common == '0' && most) || (common == '1' && !most) {
			inp = buildNewList(inp, '0', i)
		} else {
			inp = buildNewList(inp, '1', i)
		}
	}
	return inp[0]
}

func buildNewList(inp []string, b byte, pos int) []string {
	res := []string{}
	for _, v := range inp {
		if v[pos] == b {
			res = append(res, v)
		}
	}
	return res
}

func getInputs() []string {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var res []string

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}
