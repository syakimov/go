package main

import (
	"fmt"
)

type Person struct {
	FirstName, LastName string
	Money               int
}

func (p Person) Greet() {
	fmt.Println("Hello I am", p.FirstName, p.LastName)
}

func (p *Person) Spend(money int) int {
	p.Money -= money
	return p.Money
}

func (p Person) Borrow(money int) int {
	p.Money -= money
	return p.Money
}

func main() {
	p := Person{"Stef", "Yak", 20}
	p.Greet()

	fmt.Println("Money left", p.Money)
	p.Spend(5)

	fmt.Println("By Ref", p.Money)

	p.Borrow(10)
	fmt.Println("By Val", p.Money)
}
