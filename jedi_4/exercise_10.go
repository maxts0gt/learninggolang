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
	
	x["maximo_stimo"] = []string{`Super duper smart`, `Coffee`, `Fia`}
	x["spy_movie"] = []string{`Tiko smart`, `Coffee`, `Fia`}
	delete(x, "maximo_stimo")
	for i, val := range x {
		fmt.Println("This is the record for", i)
		for j, val2 := range val {
			fmt.Println("\t", j, val2)
		}
	}
	
}
