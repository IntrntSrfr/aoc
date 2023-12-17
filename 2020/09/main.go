package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"sort"
	"strings"
	"time"
)

func main() {
	inp := getInputs()

	start := time.Now()
	ans1 := part1(inp, 25)
	fmt.Println("Part 1")
	fmt.Println("Time taken:", time.Since(start))
	fmt.Println("Answer:", ans1)

	start = time.Now()
	ans2 := part2(inp, ans1)
	fmt.Println("Part 2")
	fmt.Println("Time taken:", time.Since(start))
	fmt.Println("Answer:", ans2)
}

func part1(inp []*big.Int, preamble int) *big.Int {
	for i := preamble; i < len(inp); i++ {
		if !twoSum(inp[i], copy(inp[i-preamble:i])) {
			return inp[i]
		}
	}
	return big.NewInt(-1)
}
func twoSum(target *big.Int, l []*big.Int) bool {
	sort.Sort(SortBigInt(l))
	lhs, rhs := 0, len(l)-1
	for lhs < rhs {
		sum := big.NewInt(0).Add(l[lhs], l[rhs])
		if sum.Cmp(target) == 0 {
			return true
		} else if sum.Cmp(target) < 0 {
			lhs++
		} else {
			rhs--
		}
	}
	return false
}
func copy(l []*big.Int) []*big.Int {
	nl := make([]*big.Int, len(l))
	for i, n := range l {
		m := big.NewInt(0)
		nl[i] = m.Add(m, n)
	}
	return nl
}

func part2(inp []*big.Int, target *big.Int) *big.Int {
	var l []*big.Int
	total := big.NewInt(0)
	for _, num := range inp {
		total.Add(total, num)
		l = append(l, num)
		for total.Cmp(target) > 0 {
			total.Sub(total, l[0])
			l = l[1:]
		}
		if total.Cmp(target) == 0 && len(l) > 1 {
			sort.Sort(SortBigInt(l))
			res := big.NewInt(0).Add(l[0], l[len(l)-1])
			return res
		}
	}
	return big.NewInt(-1)
}

type SortBigInt []*big.Int

func (a SortBigInt) Len() int           { return len(a) }
func (a SortBigInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBigInt) Less(i, j int) bool { return a[i].Cmp(a[j]) < 0 }

func getInputs() []*big.Int {
	data, _ := ioutil.ReadFile("./input.txt")
	var res []*big.Int
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		i := big.NewInt(0)
		i.SetString(l, 10)
		res = append(res, i)
	}
	return res
}
