package main

import "fmt"

type person struct {
	first string
}

//Interface is basically "animal being anything but can speak"
type human interface {
	speak()
}

//Defining function speak with receiver of pointer
func (p *person) speak() {
	fmt.Println("Hello")
}

//Basically speak() function comes from human interface.
func saySomething(h human) {
	h.speak()
}

func main() {
// Person whose first name is James
	p1 := person{
		first: "James",
	}


	saySomething(&p1)


// You can use below line directly. It does exactly same thing as the above one.
// Because person is pointer receiver in 'speak' function.
	p1.speak()
}
