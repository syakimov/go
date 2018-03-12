package main

import "fmt"

func main() {
	s := make([]int, 5)

	printSlice(s)
	fmt.Printf("in after ", &s[0])
}

func printSlice(s []int) {
	s[2] = 5
	fmt.Println("in     ", &s[0])
}
