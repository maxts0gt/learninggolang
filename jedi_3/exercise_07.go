package main

import (
	"fmt"
)

func main() {
	i := 4
	if i == 1 {
	fmt.Printf("I ask %v, not one", i)
	} else if i == 2 {
	fmt.Printf("I ask %v, not two", i)
	} else if i == 3 {
	fmt.Printf("Yes, this is what I asked, a number %v", i)
	} else {
	fmt.Printf("My asked number has not been found. I asked %v", i)
	}
}
