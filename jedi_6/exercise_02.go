package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	t := foo(s...)
	fmt.Println(t)
	u := []int{1, 2, 3, 4, 5}
	y := bar(u)
	fmt.Println(y)
}

func foo(s ...int) int {
	y := 0
	for _, v := range s {
		y += v

	}
	return y
}

func bar(x []int) int {
	y := 0
	for _, i := range x {
		y += i
	}
	return y
}
