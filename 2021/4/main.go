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
	fmt.Println("part 1:", finished[0].unmarked()*finished[0].LastNum)
	fmt.Println("part 2:", finished[len(finished)-1].unmarked()*finished[len(finished)-1].LastNum)
}

func solveBoards(numbers []int, boards []*Board) []*Board {
	finished := []*Board{}
	for _, num := range numbers {
		for _, board := range boards {
			if board.Finished {
				continue
			}
			board.mark(num)
			if board.bingo() {
				board.Finished = true
				board.LastNum = num
				finished = append(finished, board)
			}
		}
		if len(finished) == len(boards) {
			break
		}
	}
	return finished
}

type Board struct {
	LastNum  int
	Finished bool
	Cells    [][]*Cell
}

func newBoard() *Board {
	outer := make([][]*Cell, 5)
	for i := range outer {
		outer[i] = make([]*Cell, 5)
	}
	return &Board{Cells: outer}
}

func (b *Board) bingo() bool {
	// check if any horizontal line has bingo
	for i := range b.Cells {
		count := 0
		for j := range b.Cells[i] {
			cell := b.Cells[i][j]
			if cell.Marked {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}

	// check if any vertical line has bingo
	for i := 0; i < 5; i++ {
		count := 0
		for j := 0; j < 5; j++ {
			cell := b.Cells[j][i]
			if cell.Marked {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}

	return false
}

func (b *Board) mark(num int) {
	for i := range b.Cells {
		for j := range b.Cells[i] {
			cell := b.Cells[i][j]
			if cell.Value == num {
				cell.Marked = true
			}
		}
	}
}

func (b *Board) unmarked() int {
	sum := 0
	for i := range b.Cells {
		for j := range b.Cells[i] {
			cell := b.Cells[i][j]
			if !cell.Marked {
				sum += cell.Value
			}
		}
	}
	return sum
}

type Cell struct {
	Value  int
	Marked bool
}

func getInputs() ([]int, []*Board) {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var numbers []int

	// first scan for numbers
	scanner.Scan()
	for _, n := range strings.Split(scanner.Text(), ",") {
		nn, _ := strconv.Atoi(n)
		numbers = append(numbers, nn)
	}

	var boards []*Board

	// scan once for empty line
	for scanner.Scan() {
		b := newBoard()

		// then scan each line with numbers
		for i := 0; i < 5; i++ {
			scanner.Scan()
			f := strings.Fields(scanner.Text())
			for j, ff := range f {
				n, _ := strconv.Atoi(ff)
				b.Cells[i][j] = &Cell{Value: n}
			}
		}
		boards = append(boards, b)
	}
	return numbers, boards
}
