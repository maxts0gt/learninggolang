package main

import (
	"fmt"
)

func main() {
	switch {
	case true: fmt.Println("This seems to be true")
	case false: fmt.Println("This seems to be true but this line and belows line are not going to be printed")
	case 1 < 3: fmt.Println("It seems statement is correct")
	case 3 > 1: fmt.Println("This also seems to be true")
}
}