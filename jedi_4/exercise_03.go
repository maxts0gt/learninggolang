package main

import (
	"fmt"
)

func main() {
	
	x := []int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
	x = x[:5]
	y := x[5:10]
	z := x[2:7]
	a := x[1:6]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	
}
