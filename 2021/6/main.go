package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	state := getInputs()

	// 9999999

	start := time.Now()
	iters := 10
	for i := 0; i < iters; i++ {
		solve(deepCopy(state), 256)
	}

	fmt.Println(fmt.Sprintf("done in %v total; %v average", time.Since(start), time.Since(start)/time.Duration(iters)))
}

func solve(state []*big.Int, days int) {
	for day := 0; day < days; day++ {
		state[(day+7)%9].Add(state[(day+7)%9], state[day%9])
		//if day == 79 {
		//	fmt.Println(sum(state))
		//}
	}
	//fmt.Println(sum(state))
}

func deepCopy(s []*big.Int) []*big.Int {
	res := make([]*big.Int, len(s))
	for i := range res {
		res[i] = big.NewInt(0).Set(s[i])
	}
	return res
}

func getInputs() []*big.Int {
	f, _ := os.Open("./input.txt")
	line, _ := io.ReadAll(f)
	res := make([]*big.Int, 9)
	for i := range res {
		res[i] = big.NewInt(0)
	}
	for _, numStr := range strings.Split(string(line), ",") {
		num, _ := strconv.ParseInt(numStr, 10, 64)
		res[num].Add(res[num], big.NewInt(1))
	}
	return res
}

func sum(m []*big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, v := range m {
		sum.Add(sum, v)
	}
	return sum
}
