package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

}

func getInputs() []int {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var res []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		res = append(res, n)
	}

	return res
}
