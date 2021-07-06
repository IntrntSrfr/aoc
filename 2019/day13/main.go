package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	inp := getList()

	b := 0

	for i, v := range inp {
		if i%2 == 0 && v == "2" {
			b++
		}
	}
	fmt.Println(b)

}

func getList() []string {

	f, _ := os.Open("./input.txt")
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	v := strings.Split(string(b), ",")

	return v
}
