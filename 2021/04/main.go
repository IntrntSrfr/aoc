package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var boardSize = 5

func main() {
	parserStart := time.Now()
	numbers, boards := getInputs()
	fmt.Println("time taken to parse:", time.Since(parserStart))

	start := time.Now()
	finished := solveBoards(numbers, boards)
	fmt.Println("time taken to solve all boards:", time.Since(start))

	fmt.Println("part 1:", finished[0])
	fmt.Println("part 2:", finished[len(finished)-1])
}

func solveBoards(numbers []int, boards []*Board) []int {
	finished := []int{}
	for _, num := range numbers {
		for _, board := range boards {
			if board.Finished {
				continue
			}
			board.mark(num)
			if board.bingo() {
				board.Finished = true
				finished = append(finished, board.UnmarkedSum*num)
			}
		}
		if len(finished) == len(boards) {
			break
		}
	}
	return finished
}

type Board struct {
	Finished    bool
	Cells       [][]int
	RowCount    []int
	ColCount    []int
	UnmarkedSum int
}

func newBoard() *Board {
	outer := make([][]int, boardSize*boardSize)
	return &Board{Cells: outer, RowCount: make([]int, boardSize), ColCount: make([]int, boardSize)}
}

func (b *Board) bingo() bool {
	for i := 0; i < boardSize; i++ {
		if b.RowCount[i] == boardSize || b.ColCount[i] == boardSize {
			return true
		}
	}
	return false
}

func (b *Board) mark(num int) {
	for i := range b.Cells {
		cell := b.Cells[i]
		if cell[0] == num && cell[1] == 0 {
			b.UnmarkedSum -= num
			b.RowCount[i/boardSize]++
			b.ColCount[i%boardSize]++
			cell[1] = 1
		}
	}
}

func getInputs() ([]int, []*Board) {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var numbers []int
	scanner.Scan()
	for _, n := range strings.Split(scanner.Text(), ",") {
		nn, _ := strconv.Atoi(n)
		numbers = append(numbers, nn)
	}

	var boards []*Board
	for scanner.Scan() {
		b := newBoard()
		for i := 0; i < boardSize; i++ {
			scanner.Scan()
			f := strings.Fields(scanner.Text())
			for j, ff := range f {
				n, _ := strconv.Atoi(ff)
				b.Cells[i*boardSize+j] = []int{n, 0}
				b.UnmarkedSum += n
			}
		}
		boards = append(boards, b)
	}
	return numbers, boards
}
