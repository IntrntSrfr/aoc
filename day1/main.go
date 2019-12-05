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
	for scanner.Scan() {
		n, _ := strconv.ParseFloat(scanner.Text(), 64)
		total += calculateFuel2(n)

	}

	fmt.Println(total)

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
