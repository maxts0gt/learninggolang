package main

import (
	"fmt"
)

type person struct {
	first     string
	last      string
	favFlavor []string
}

func main() {

	p1 := person{
		first:     "Max",
		last:      "Tsogi",
		favFlavor: []string{"chocolate", "mint-chocolate"},
	}

	p2 := person{
		first:     "Emma",
		last:      "Eki",
		favFlavor: []string{"vanilla", "yoghurt"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k, v.first)
		for i, val := range v.favFlavor {
			fmt.Println(i, val)
		}

	}
}
