package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	t, list := getInputs()

	lowestDiff, lowestID := 10000, 10000
	for i := 0; i < len(list); i++ {
		n, err := strconv.Atoi(list[i])
		if err != nil {
			continue
		}
		m := t % n
		diff := n - m
		if diff < lowestDiff {
			lowestDiff = diff
			lowestID = n
		}
	}
	fmt.Println(lowestID, lowestDiff, lowestID*lowestDiff)
}

func getInputs() (int, []string) {
	data, _ := ioutil.ReadFile("./input.txt")
	parts := strings.Split(string(data), "\n")
	n, _ := strconv.Atoi(parts[0])
	lines := strings.Split(parts[1], ",")
	return n, lines
}
