package main

import (
	fmt "fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inp := getLines()

	part1("shiny gold", inp)

}

func part1(end string, inp []string) {
	for _, line := range inp {
		i := strings.Split(line, " bags contain ")
		j := strings.Split(i[1], ", ")
		if in(j, end) && !in(bags, i[0]) {
			bags = append(bags, i[0])
			part1(i[0], inp)
		}
	}
	fmt.Println(len(bags))
}

var bags []string

func search(end string, m map[string][]string) {
	for k, v := range m {
		if in(v, end) && !in(bags, k) {
			bags = append(bags, k)
			search(k, m)
		}
	}
}

func in(l []string, v string) bool {
	for _, w := range l {
		return strings.Contains(w, v)
	}
	return false
}

func getLines() []string {
	data, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(data), "\n")
}
