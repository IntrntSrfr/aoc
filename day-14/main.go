package main

import "fmt"

func main() {
	fmt.Println(fmt.Sprintf("%010d", 50))

	mask := make([]string, 36)
	for i := range mask {
		mask[i] = "X"
	}

	fmt.Println(mask)

}
