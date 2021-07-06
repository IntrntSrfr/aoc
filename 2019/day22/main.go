package main

import (
	"fmt"
)

func main() {
	//deck := [10007]int{}
	deck := []int{}
	for i := 0; i < 10; i++ {
		deck = append(deck, i)
	}
	/*
		fmt.Println(deck)
		sort.Sort(sort.Reverse(sort.IntSlice(deck)))
		fmt.Println(deck)
	*/
	deck = cut(deck, -4)

	fmt.Println(deck)

}

func cut(i []int, n int) []int {

	if n < 0 {
		temp := i[:-n]
		fmt.Println(temp)
		i = i[:len(i)+n]
		fmt.Println(i)
		i = append(i, temp...)

	} else {
		temp := i[:n]
		i = i[n:]
		i = append(i, temp...)
	}

	return i
}
