package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("./input.txt")
	parts := strings.Split(string(data), "\n\n")

	tickets := [][]int{}
	for _, ticket := range strings.Split(parts[2], "\n") {
		var values []int
		for _, val := range strings.Split(ticket, ",") {
			n, _ := strconv.Atoi(val)
			values = append(values, n)
		}
		tickets = append(tickets, values)
	}
	tickets = tickets[1:]

	fmt.Println(tickets)

	sum := 0
	for _, field := range strings.Split(parts[0], "\n") {
		line := strings.Split(field, ": ")
		//fmt.Println(line[1])

		ranges := strings.Split(line[1], " or ")
		//fmt.Println(ranges)

		for _, ticket := range tickets {

			valid := false
			badvalue := 0

			for _, r := range ranges {
				spl := strings.Split(r, "-")
				start, _ := strconv.Atoi(spl[0])
				end, _ := strconv.Atoi(spl[1])

				//fmt.Println(start, end)

				for _, val := range ticket {
					if val < start || val > end {
						//valid = false
						//		fmt.Println(val, start, end )
						badvalue = val
						//fmt.Println(start, end, val)
						continue
					}
					valid = true
				}
			}
			if !valid {
				fmt.Println(badvalue)
				sum += badvalue
			}
		}
	}

	fmt.Println(sum)

	//fmt.Println(tickets)

}
