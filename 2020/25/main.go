package main

import (
	"fmt"
)

func main() {
	k1 := 3248366
	k2 := 4738476

	ls1 := dec(k1, 7)
	ls2 := dec(k2, 7)

	fmt.Println(ls1, ls2)
	fmt.Println(k1, k2)

	tk1 := enc(k1, ls2)
	tk2 := enc(k2, ls1)

	fmt.Println(tk1, tk2)
	if tk1 == tk2 {
		fmt.Println("found good keys")
	}
}

func dec(t, s int) int {
	ls := 0
	for ts := 1; ts != t; ls++ {
		ts *= s
		ts %= 20201227
	}
	return ls
}

func enc(s, ls int) int {
	val := 1
	for i := 0; i < ls; i++ {
		val *= s
		val %= 20201227
	}
	return val
}
