package main

// Lists does not have contains
// List Tricks: https://github.com/golang/go/wiki/SliceTricks
import "fmt"

func main() {
	s := []int{1, 2}
	fmt.Println(s, len(s))
	s = append(s, 3)
	fmt.Println(s, len(s))
}
