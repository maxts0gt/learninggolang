package main

import (
	"fmt"
)

type person struct {
	firstName   string
	lastName    string
	favIcecream []string
}

func main() {

	max := person{
		firstName:   "Maxi",
		lastName:    "Tsogi",
		favIcecream: []string{"chocolate", "mint-chocolate", "vanilla"},
	}

	emma := person{
		firstName:   "Emma",
		lastName:    "Eki",
		favIcecream: []string{"vanilla", "mint-chocolate", "yoghurt"},
	}

	fmt.Println(max.firstName, max.lastName)
	for i, v := range max.favIcecream {
		fmt.Println(i, v)
	}
	fmt.Println(emma.firstName, emma.lastName)
	for i, v := range emma.favIcecream {
		fmt.Println(i, v)
	}

}
