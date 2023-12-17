package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(getSeq())
}

func getSeq() []int {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	v := strings.Split(string(b), "")

	res := make([]int, len(v))
	for i := range v {
		n, _ := strconv.Atoi(v[i])
		res[i] = n
	}
	return res
}
