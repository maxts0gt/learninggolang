package main

import (
	"fmt"
)

func main() {
	favSport := "Basketball"
	switch favSport {
	case "Baseball":
		fmt.Println("This seems to be true")
	case "Football":
		fmt.Println("This seems to be true but this line and belows line are not going to be printed")
	case "Volleyball":
		fmt.Println("It seems statement is correct")
	case "Basketball":
		fmt.Printf("This is the only that I edited. Case seems to be %v", favSport)
	}
}
