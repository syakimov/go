package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	// Create
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// Access field
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
}
