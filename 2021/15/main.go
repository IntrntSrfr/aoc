package main

import (
	"bufio"
	"fmt"
	"github.com/intrntsrfr/aoc/utils"
	"os"
)

//var moves = [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
var moves = [][]int{{0, 1}, {1, 0}}

func main() {
	grid := getInputs()
	shortestPathScore(grid, []int{0, 0}, []int{99, 99})

	newGrid := tileGrid(grid)
	//utils.DisplayGrid(newGrid)
	shortestPathScore(newGrid, []int{0, 0}, []int{499, 499})
}

func tileGrid(grid [][]int) [][]int {
	newGrid := utils.MxNGrid(len(grid)*5, len(grid)*5)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for y := range grid {
				for x := range grid {
					val := grid[y][x] + i + j
					if (grid[y][x] + i + j) > 9 {
						val = ((grid[y][x] + i + j) % 10) + 1
					}
					newGrid[y+len(grid)*i][x+len(grid)*j] = val
				}
			}
		}
	}
	return newGrid
}

func shortestPathScore(grid [][]int, start, end []int) {
	var visited = make(map[string]bool)
	var predecessors = make(map[string][]int)
	var cellCosts = make(map[string]int)

	pq := NewPrioQueue()

	pq.Enqueue(&Node{Data: start, Prio: 0})
	predecessors[makeKey(start[0], start[1])] = start
	visited[makeKey(start[0], start[1])] = true

	for !pq.Empty() {
		n := pq.Dequeue()
		curY, curX := n.Data[0], n.Data[1]
		cur := n.Data
		if curY == len(grid)-1 && curX == len(grid[0])-1 {
			break
		}
		for _, move := range moves {
			newY, newX := curY+move[0], curX+move[1]
			if newY < 0 || newX < 0 || newY > len(grid)-1 || newX > len(grid)-1 {
				continue
			}
			key := makeKey(newY, newX)
			newCost := grid[newY][newX] + n.Prio // + MDistance(newY, newX, end[0], end[1])
			if !visited[key] || newCost < cellCosts[key] {
				pq.Enqueue(&Node{Data: []int{newY, newX}, Prio: newCost})
				predecessors[key] = cur
				cellCosts[key] = newCost
				visited[key] = true
			}
		}
	}

	score := 0
	cur := end
	for makeKey(cur[0], cur[1]) != "0-0" {
		score += grid[cur[0]][cur[1]]
		cur = predecessors[makeKey(cur[0], cur[1])]
	}
	fmt.Println("score", score)
}

func MDistance(y0, x0, y1, x1 int) int {
	return utils.Abs(x0-x1) + utils.Abs(y0-y1)
}

func makeKey(y, x int) string {
	return fmt.Sprintf("%v-%v", y, x)
}

func getInputs() [][]int {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var res [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var lineRes []int
		for _, c := range line {
			lineRes = append(lineRes, int(c)-48)
		}
		res = append(res, lineRes)
	}
	return res
}

type PrioQueue struct {
	Head *Node
}

func NewPrioQueue() *PrioQueue {
	return &PrioQueue{Head: nil}
}

func (p *PrioQueue) Empty() bool {
	return p.Head == nil
}

func (p *PrioQueue) Enqueue(n *Node) {
	if p.Empty() {
		p.Head = n
		return
	}
	if n.Prio < p.Head.Prio {
		n.Next = p.Head
		p.Head = n
		return
	}
	cur := p.Head
	for ; cur.Next != nil; cur = cur.Next {
		if n.Prio < cur.Next.Prio {
			n.Next = cur.Next
			cur.Next = n
			return
		}

	}
	cur.Next = n
}

func (p *PrioQueue) Dequeue() *Node {
	if p.Empty() {
		return nil
	}
	h := p.Head
	p.Head = h.Next
	h.Next = nil
	return h
}

func (p *PrioQueue) Display() {
	fmt.Println("display:")
	cur := p.Head
	for ; cur != nil; cur = cur.Next {
		fmt.Println(cur)
	}
}

type Node struct {
	Prio int
	Data []int
	Next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("prio: %v; data: %v", n.Prio, n.Data)
}
