package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	f, _ := os.Open("./input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0.0
	total2 := 0.0
	for scanner.Scan() {
		n, _ := strconv.ParseFloat(scanner.Text(), 64)
		total += calculateFuel(n)
		total2 += calculateFuel2(n)
	}

	fmt.Println("part 1:", total)
	fmt.Println("part 2:", total2)

}

func calculateFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func calculateFuel2(mass float64) float64 {
	total := 0.0
	add := calculateFuel(mass)

	for add > 0 {
		total += add
		add = calculateFuel(add)
	}
	return total
}
