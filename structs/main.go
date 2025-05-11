package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex.updateName("Jim")
	alex.print()
}

func (p *person) updateName(n string){
	p.firstName = n
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
