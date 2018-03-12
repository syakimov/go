package main

import "fmt"

func main() {
	m := make(map[string]int)
	printMap(m)
	fmt.Println("after function %s", m)

}

func printMap(m map[string]int) {
	m["five"] = 5
	fmt.Printf("after", m[0])
}
