package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "M",
		Age:   29,
	}

	u2 := User{
		First: "E",
		Age:   28,
	}

	u3 := User{
		First: "O",
		Age:   5,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
