package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	passes := getInputs()

	start := time.Now()

	highest := 0
	var list []int
	for _, pass := range passes {
		pass = strings.ReplaceAll(pass, "F", "0")
		pass = strings.ReplaceAll(pass, "B", "1")
		pass = strings.ReplaceAll(pass, "L", "0")
		pass = strings.ReplaceAll(pass, "R", "1")

		row, _ := strconv.ParseInt(pass[:7], 2, 64)
		seat, _ := strconv.ParseInt(pass[len(pass)-3:], 2, 64)

		sID := int(row*8 + seat)
		if sID > highest {
			highest = sID
		}
		list = append(list, sID)
	}
	fmt.Println(highest)

	for i := 0; i < highest; i++ {
		if inList(list, i-1) && inList(list, i+1) && !inList(list, i) {
			fmt.Println(i)
		}
	}

	fmt.Println(time.Since(start))

	/*
		sort.Ints(list)
		for i := 1; i < len(list)-1; i++ {
			if list[i+1] != list[i]+1 {
				fmt.Println(list[i]+1)
			}
		}
	*/
}

func inList(list []int, n int) bool {
	for _, l := range list {
		if l == n {
			return true
		}
	}
	return false
}

func getInputs() []string {
	data, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(data), "\n")
}
