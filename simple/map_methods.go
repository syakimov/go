package main

import "fmt"

func main() {
	// map[name]age
	people := map[string]int{"Stefan": 20}

	// Add
	people["Pepi"] = 19

	// Get
	age, present := people["Stefan"]

	if !present {
		fmt.Println("Handle error!!")
	}

	// Delete
	delete(people, "Pepi")
	_, present = people["Pepi"]

	if present {
		fmt.Println("Handle error!!")
	}

	people["Moni"] = age + 10

	fmt.Println("--------------------------------")
	fmt.Println(people)
	fmt.Println("--------------------------------")
}
