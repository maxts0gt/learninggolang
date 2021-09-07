package main

import (
	"fmt"
)

func main() {
	f := struct {
		first   string
		friends map[string]string
		looks   []string
	}{
		first: "Joey",
		friends: map[string]string{
			"Friend1": "Chandler",
			"Friend2": "Ross",
		},
		looks: []string{"funny", "tall"},
	}



	for k, v := range f.friends {
		fmt.Println(k, v)
		for k, v := range f.looks {
			fmt.Println(k, v)
		}
	}
}
