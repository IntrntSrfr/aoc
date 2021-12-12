package utils

import (
	"fmt"
	"math/big"
	"sort"
)

func DisplayGrid(g [][]int) {
	fmt.Println("grid display:")
	for _, v := range g {
		fmt.Println(v)
	}
}

func PadStringGrid() {

}

type BigIntSlice []*big.Int

func (s BigIntSlice) Len() int           { return len(s) }
func (s BigIntSlice) Less(i, j int) bool { return s[i].Cmp(s[j]) < 0 }
func (s BigIntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Sort is a convenience method.
func (s BigIntSlice) Sort() {
	sort.Sort(s)
}

func Sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func Product(a []int) int {
	product := 1
	for _, v := range a {
		product *= v
	}
	return product
}

func SumNthPos(a [][]int, pos int) int {
	sum := 0
	for _, v := range a {
		sum += v[pos]
	}
	return sum
}

func ProductNthPos(a [][]int, pos int) int {
	product := 1
	for _, v := range a {
		product *= v[pos]
	}
	return product
}

type Stack struct {
	top    *Node
	length int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Empty() bool {
	return s.length == 0
}

func (s *Stack) Peek() int {
	if s.Empty() {
		return -1
	}
	return s.top.value
}

func (s *Stack) Pop() int {
	if s.Empty() {
		return -1
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(value int) {
	n := &Node{value, s.top}
	s.top = n
	s.length++
}

type Node struct {
	value int
	prev  *Node
}

type Queue struct {
	data [][]int
}

func NewQueue() *Queue {
	return &Queue{
		data: [][]int{},
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Push(v []int) {
	q.data = append(q.data, v)
}

func (q *Queue) Pop() []int {
	if q.IsEmpty() {
		return []int{}
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
