package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(intersect("abcdfg", "abcdefg"))
	fmt.Println(complement("abcdfg", "abcdefg"))

	/*
		inp := getInputs()
		//count := 0
		for _, line := range inp {
			//mapping := make(map[string]string)
			fmt.Println(line[:10])
				for _, code := range line[:10] {
					fmt.Println(code)
				}
			}
	*/
}

func intersect(a, b string) string {
	res := ""
	for _, c1 := range a {
		if strings.Contains(b, string(c1)) {
			res += string(c1)
		}
	}
	return res
}

func complement(a, b string) string {
	str := "abcdef"
	res := ""

	isct := intersect(a, b)

	for _, c1 := range str {
		if !strings.Contains(isct, string(c1)) {
			res += string(c1)
		}
	}

	return res
}

func part1() {
	inp := getInputs()
	count := 0
	for _, line := range inp {
		for _, code := range line[11:] {
			n := getDigit(code)
			if n == 1 || n == 7 || n == 4 || n == 8 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func getDigit(encoding string) int {
	switch len(encoding) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	case 5:
		// nums can be 5, 2 or 3
		if strings.Contains(encoding, "a") && strings.Contains(encoding, "d") && strings.Contains(encoding, "g") {
			// strings must contain a, d or g
			if strings.Contains(encoding, "c") && strings.Contains(encoding, "e") {
				return 2
			} else if strings.Contains(encoding, "c") && strings.Contains(encoding, "f") {
				return 3
			} else if strings.Contains(encoding, "b") && strings.Contains(encoding, "f") {
				return 5
			}
		}
	case 6:
		// nums can be 0, 6 or 9
		if strings.Contains(encoding, "a") && strings.Contains(encoding, "b") && strings.Contains(encoding, "f") && strings.Contains(encoding, "g") {
			// strings must contain a, b, f or g
			if strings.Contains(encoding, "c") && strings.Contains(encoding, "e") {
				return 0
			} else if strings.Contains(encoding, "d") && strings.Contains(encoding, "e") {
				return 6
			} else if strings.Contains(encoding, "c") && strings.Contains(encoding, "d") {
				return 9
			}
		}
	}
	return -1
}

func getInputs() [][]string {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines [][]string
	for scanner.Scan() {
		lines = append(lines, strings.Fields(scanner.Text()))
	}
	return lines
}
