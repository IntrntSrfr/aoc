package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp := getInputs()
	part1(inp)
}

func part2(inp [][]string) {

	mem := make(map[int]int64)
	mask := ""
	for _, l := range inp {
		if l[0] == "mask" {
			mask = l[1]
		} else if strings.HasPrefix(l[0], "mem[") {
			l[0] = l[0][4 : len(l[0])-1]
			index, _ := strconv.ParseInt(l[0], 10, 64)
			value, _ := strconv.Atoi(l[1])
			indexStr := []byte(fmt.Sprintf("%036b", index))
			for i := range indexStr {
				if mask[i] == '0' {
					continue
				} else if mask[i] == '1' {
					indexStr[i] = '1'
				} else {

				}

			}

			value, _ = strconv.ParseInt(string(valueStr), 2, 64)
			mem[index] = value

		}
	}

	sum := int64(0)
	for _, v := range mem {
		sum += v
	}

	fmt.Println(sum)
}

func part1(inp [][]string) {

	mem := make(map[int]int64)
	mask := ""
	for _, l := range inp {
		if l[0] == "mask" {
			mask = l[1]
		} else if strings.HasPrefix(l[0], "mem[") {
			l[0] = l[0][4 : len(l[0])-1]
			index, _ := strconv.Atoi(l[0])
			value, _ := strconv.ParseInt(l[1], 10, 64)
			valueStr := []byte(fmt.Sprintf("%036b", value))
			for i := range valueStr {
				if mask[i] == 'X' {
					continue
				}
				valueStr[i] = mask[i]
			}
			value, _ = strconv.ParseInt(string(valueStr), 2, 64)
			mem[index] = value
		}
	}

	sum := int64(0)
	for _, v := range mem {
		sum += v
	}

	fmt.Println(sum)
}

func getInputs() [][]string {
	var res [][]string

	f, _ := os.Open("./input.txt")
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		l := scan.Text()
		res = append(res, strings.Split(l, " = "))
	}
	return res
}
