package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	instructions := getInstructions()
	start := time.Now()
	part1(instructions)
	fmt.Println("Part 1 time:", time.Since(start))
	start = time.Now()
	part2(instructions)
	fmt.Println("Part 2 time:", time.Since(start))
}

func part1(inp []Instruction) {
	ans, _ := runInstructions(inp)
	fmt.Println(ans)
}
func part2(inp []Instruction) {
	for i := 0; i < len(inp); i++ {
		oldOp := inp[i].operation
		switch inp[i].operation {
		case "nop":
			inp[i].operation = "jmp"
		case "jmp":
			inp[i].operation = "nop"
		default:
			continue
		}
		ans, ok := runInstructions(inp)
		if ok {
			fmt.Println(ans)
			return
		}
		inp[i].operation = oldOp
	}
}
func runInstructions(inp []Instruction) (int, bool) {
	acc, index, seen := 0, 0, make(map[int]bool)
	for index < len(inp) && index >= 0 && !seen[index] {
		seen[index] = true
		inst := inp[index]
		switch inst.operation {
		case "nop":
			index++
		case "acc":
			acc += inst.value
			index++
		case "jmp":
			index += inst.value
		}
	}
	return acc, index >= len(inp)
}
func getInstructions() []Instruction {
	data, _ := ioutil.ReadFile("./huge.txt")
	lines := strings.Split(string(data), "\n")
	var res []Instruction
	for _, line := range lines {
		inst := strings.Split(line, " ")
		n, _ := strconv.Atoi(inst[1])
		res = append(res, Instruction{operation: inst[0], value: n})
	}
	return res
}

type Instruction struct {
	operation string
	value     int
}
