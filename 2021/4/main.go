package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers, boards := getInputs()
	finished := solveBoards(numbers, boards)
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
				finished = append(finished, board.unmarkedSum()*num)
			}
		}
		if len(finished) == len(boards) {
			break
		}
	}
	return finished
}

type Board struct {
	Finished bool
	Cells    [][]int
	RowCount []int
	ColCount []int
}

func newBoard() *Board {
	outer := make([][]int, 25)
	return &Board{Cells: outer, RowCount: make([]int, 5), ColCount: make([]int, 5)}
}

func (b *Board) bingo() bool {
	for i := 0; i < 5; i++ {
		if b.RowCount[i] == 5 || b.ColCount[i] == 5 {
			return true
		}
	}
	return false
}

func (b *Board) mark(num int) {
	for i := range b.Cells {
		cell := b.Cells[i]
		if cell[0] == num && cell[1] == 0 {
			b.RowCount[i/5]++
			b.ColCount[i%5]++
			cell[1] = 1
		}
	}
}

func (b *Board) unmarkedSum() int {
	sum := 0
	for i := range b.Cells {
		cell := b.Cells[i]
		if cell[1] == 0 {
			sum += cell[0]
		}
	}
	return sum
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
		for i := 0; i < 5; i++ {
			scanner.Scan()
			f := strings.Fields(scanner.Text())
			for j, ff := range f {
				n, _ := strconv.Atoi(ff)
				b.Cells[i*5+j] = []int{n, 0}
			}
		}
		boards = append(boards, b)
	}
	return numbers, boards
}


