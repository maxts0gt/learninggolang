package main



import (
	"fmt"
)



func main() {
	
	defer foo()
	bar()
	fmt.Println("Hello, playground")
}

func foo() {
	fmt.Println("This should be the first line")
}

func bar() {
	fmt.Println("This is actaully the second line")
}