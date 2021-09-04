package main

import "fmt"

var number = 42


func main() {
	fmt.Println(number)
	fmt.Printf("%d\n%b\n%#x", number, number, number)
	another_number := number << 1
	fmt.Println(another_number)
	fmt.Printf("%d\n%b\n%#x", another_number, another_number, another_number)
}

