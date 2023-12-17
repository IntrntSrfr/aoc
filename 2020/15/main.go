package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inp := getInputs()

	sinceSpoken := make(map[int]int)

	turn := 0

	fmt.Println(inp)
	for _, val := range inp {

		sinceSpoken[val] = 0

		for k := range sinceSpoken {
			if k != val {
				sinceSpoken[k]++
			}
		}
		turn++
	}
	fmt.Println(sinceSpoken)
	fmt.Println(turn)

	for {

	}

}

func getInputs() []int {
	data, _ := ioutil.ReadFile("./input.txt")
	var res []int
	for _, l := range strings.Split(string(data), ",") {
		n, _ := strconv.Atoi(l)
		res = append(res, n)
	}
	return res
}
