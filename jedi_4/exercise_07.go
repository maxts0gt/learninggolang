package main

import (
	"fmt"
)

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Hellooooo, James"} 
	z := [][]string{x, y}
	fmt.Println(z)
}
