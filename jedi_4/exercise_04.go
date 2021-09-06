package main

import (
	"fmt"
)

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
	x = append(x, 53)
	fmt.Println(x)
	x = append(x, 54, 55, 56)
	fmt.Println(x)
	y := []int{57, 58, 59, 60, 61}
	x = append(x, y...)
	fmt.Println(x)
}
