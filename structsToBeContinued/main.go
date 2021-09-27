package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}


func main(){
jim := person {
	firstName: "jim",
	lastName: "jumbo",
	contactInfo: contactInfo{
		email: "jim@jimbo.com",
		zipCode: 94000,
	},
}

 //jimPointer := &jim
jim.updateName("jimmy")
jim.print()
}

func (pointerToPerson person) updateName(newFirstName string){
	(pointerToPerson).firstName= newFirstName

}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}