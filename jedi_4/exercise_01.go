package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	for i := range x {
		fmt.Println(i)
	}
}
