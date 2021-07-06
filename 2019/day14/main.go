package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	needs := make(map[string]int)
	recipes := getRecipes()

	needs["FUEL"] = 1

	fmt.Println(ore(map[string]int{"FUEL": 1}, recipes))

	fmt.Println(sort.Search(100000, func(n int) bool {
		return ore(map[string]int{"FUEL": n}, recipes) > 100000
	}) - 1)
	/*
		amt := 1000000000000
		f := 0
		for amt > 0 {
	*/

	/*
			amt -= needs["ORE"]
			f++
		}

		sort.Search()

		fmt.Println(f)
	*/
	//fmt.Println(recipes)
}
func ore(needs map[string]int, recipes map[string]recipe) int {
	ok := true
	for ok {
		ok = false

		for i, v := range needs {
			if i == "ORE" {
				continue
			}
			if v <= 0 {
				continue
			} else if v > 0 {
				ok = true
			}

			for j, v2 := range recipes[i].input {
				needs[j] += v2
			}
			needs[i] -= recipes[i].amount
		}
	}
	return needs["ORE"]
}

func getRecipes() map[string]recipe {

	r := map[string]recipe{}

	f, _ := os.Open("./input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " => ")

		// out
		out := strings.Split(s[len(s)-1], " ")
		n, _ := strconv.Atoi(out[0])

		// inp
		inpm := make(map[string]int)
		inps := strings.Split(s[0], ", ")
		for _, s2 := range inps {
			s3 := strings.Split(s2, " ")
			n, _ := strconv.Atoi(s3[0])
			inpm[s3[1]] = n
		}
		r[out[1]] = recipe{amount: n, input: inpm}
	}

	return r
}

type recipe struct {
	input  map[string]int
	amount int
}
