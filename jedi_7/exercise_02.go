package main

import (
	"fmt"
)

type person struct {
	f string
}

func main() {
	p1 := person{
		f: "old name",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}

func changeMe(a *person) {
	a.f = "lazy Name"
}
