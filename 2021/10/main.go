package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/intrntsrfr/aoc/utils"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	p1Score := 0
	p2Scores := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		valid, s := isValid(line)
		if !valid {
			p1Score += s
		} else {
			p2Scores = append(p2Scores, s)
		}
	}
	sort.Ints(p2Scores)
	fmt.Println("part 1 answer:", p1Score)
	fmt.Println("part 2 answer:", p2Scores[len(p2Scores)/2])
}

func isValid(s string) (bool, int) {
	st := utils.NewStack()
	for _, c := range s {
		if c == '{' || c == '(' || c == '[' || c == '<' {
			st.Push(int(c))
			continue
		}
		if st.Peek() != p1map[c][0] {
			return false, p1map[c][1]
		}
		st.Pop()
	}
	score := 0
	for !st.Empty() {
		score = score*5 + p2map[st.Pop()]
	}
	return true, score
}

var p1map = map[rune][]int{
	')': {int('('), 3},
	']': {int('['), 57},
	'}': {int('{'), 1197},
	'>': {int('<'), 25137},
}
var p2map = map[int]int{int('('): 1, int('['): 2, int('{'): 3, int('<'): 4}
