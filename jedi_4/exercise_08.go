package main

import (
	"fmt"
)

func main() {

	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for i, val := range x {
		fmt.Println("This is the record for", i)
		for j, val2 := range val {
			fmt.Println("\t", j, val2)
		}
	}
}
