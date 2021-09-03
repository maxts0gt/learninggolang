package main

import "fmt"

type mytype int

var x mytype

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println("\n", x)
}