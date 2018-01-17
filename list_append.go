package main

// Lists does not have contains
import "fmt"

func main() {
	s := []int{1, 2}
	fmt.Println(s, len(s))
	s = append(s, 3)
	fmt.Println(s, len(s))
}
