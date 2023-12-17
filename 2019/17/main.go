package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	inp := getList()

	

}

func getList() []string {

	f, _ := os.Open("./input.txt")
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	v := strings.Split(string(b), ",")

	return v
}
