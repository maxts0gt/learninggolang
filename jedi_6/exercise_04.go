package main

import (
	"fmt"
)

type person struct {
	first, last string
	age         int
}

func (p person) talker() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	r := person{
		first: "Max",
		last:  "Yiukae",
		age:   22,
	}
	r.talker()
}
