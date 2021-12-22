package main

import (
	"bufio"
	"os"
)

func main() {

}

func getInputs() []string {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var res []string
	for scanner.Scan() {

	}

	return res
}
