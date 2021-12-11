package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"time"

	"github.com/intrntsrfr/aoc/utils"
)

func getInputs()[]string{
	f, _ := os.Open("./10_7000_10000.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	inp := getInputs()
	
	start := time.Now()
	p1Score := big.NewInt(0)
	p2Scores := []*big.Int{}
	for _, line:=range inp {
		valid, s := isValid(line)
		if !valid {
			p1Score = p1Score.Add(p1Score, s)
		} else {
			p2Scores = append(p2Scores, s)
		}
	}
	fmt.Println("part 1 answer:", p1Score)
	sort.Sort(utils.BigIntSlice(p2Scores))
	fmt.Println("part 2 answer:", p2Scores[len(p2Scores)/2])
	fmt.Println("time taken:", time.Since(start))
}

func isValid(s string) (bool, *big.Int) {
	st := utils.NewStack()
	for _, c := range s {
		if c == '{' || c == '(' || c == '[' || c == '<' {
			st.Push(int(c))
			continue
		}
		if st.Peek() != p1map[c][0] {
			return false, big.NewInt(int64(p1map[c][1]))
		}
		st.Pop()
	}
	score := big.NewInt(0)
	for !st.Empty() {
		score = score.Mul(score, big.NewInt(5))
		score = score.Add(score, big.NewInt(int64(p2map[st.Pop()])))
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
