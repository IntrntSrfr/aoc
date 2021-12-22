package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	initial, instructions := getInputs()

	fmt.Println(initial, instructions)

	offset := 0
	for i := 0; i < 10; i++ {
		newString := initial
		for _, instruction := range instructions {
			strings.Index()

		}
	}
}

func getInputs() (string, [][]string) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	initial := scanner.Text()
	scanner.Scan()

	var instructions [][]string
	for scanner.Scan() {
		instructions = append(instructions, strings.Split(scanner.Text(), " -> "))
	}

	return initial, nil
}
