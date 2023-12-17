package main

import (
	"fmt"
	"strconv"
)

func main() {

	lower := 273025
	upper := 767253

	fmt.Println(password(lower, upper))

}

func password(lower, upper int) int {

	passes := 0

	for i := lower; i <= upper; i++ {
		m := make(map[byte]int)
		for j := 0; j < 6; j++ {
			m[strconv.Itoa(i)[j]]++
		}

		adj := false
		desc := false
		for j := 0; j < 5; j++ {
			if strconv.Itoa(i)[j] > strconv.Itoa(i)[j+1] {
				desc = true
				break
			}

			if strconv.Itoa(i)[j] == strconv.Itoa(i)[j+1] {
				adj = true
			}
		}
		if desc {
			continue
		}
		if !adj {
			continue
		}

		ok := false
		for _, v := range m {
			if v == 2 {
				ok = true
				break
			}
		}
		if !ok {
			continue
		}
		passes++
	}

	return passes
}
