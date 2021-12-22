package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	t, g := getInputs()
	fmt.Println(t)
	fmt.Println(g)
}

func getInputs() (string, [][]string) {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var res [][]string
	scanner.Scan()
	table := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		var row []string
		for _, c := range scanner.Text() {
			row = append(row, string(c))
		}
		res = append(res, row)
	}
	return table, res
}
