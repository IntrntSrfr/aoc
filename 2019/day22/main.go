package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, _ := os.Open("./input.txt")
	defer f.Close()

	deck := []int{}
	for i := 0; i < 10007; i++ {
		deck = append(deck, i)
	}

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		s := strings.Split(scan.Text(), " ")

		if strings.HasPrefix(scan.Text(), "cut") {
			n, _ := strconv.Atoi(s[1])
			deck = cut(deck, n)
		} else if strings.HasPrefix(scan.Text(), "deal with") {
			n, _ := strconv.Atoi(s[3])
			deck = increment(deck, n)
		} else if strings.HasPrefix(scan.Text(), "deal into") {
			deck = newDeck(deck)
		}
	}

	for i := range deck {
		if deck[i] == 2019 {
			fmt.Println(i)
		}
	}

}
func increment(deck []int, n int) []int {

	asd := make([]int, len(deck))
	for i := range deck {
		asd[(i*n)%len(deck)] = deck[i]
	}

	return asd

}
func newDeck(deck []int) []int {

	for i := len(deck)/2 - 1; i >= 0; i-- {
		opp := len(deck) - 1 - i
		deck[i], deck[opp] = deck[opp], deck[i]
	}
	return deck

}

func cut(i []int, n int) []int {

	if n < 0 {
		temp := i[:len(i)+n]
		i = i[len(i)+n:]
		i = append(i, temp...)
	} else {
		temp := i[:n]
		i = i[n:]
		i = append(i, temp...)
	}

	return i
}
