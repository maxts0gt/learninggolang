package main

import "fmt"

type mytype int

var x mytype
var y int


func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println("\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T",y)
}