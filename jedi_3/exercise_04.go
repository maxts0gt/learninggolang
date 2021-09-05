package main

import (
	"fmt"
)

func main() {
	i := 1990
	for {
		if i == 2022 {
			break
		}
		fmt.Println(i)
		i++
	}
}
