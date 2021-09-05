package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%v is the remainder when %v is divided by 4\n", i%4, i)
	}
}
