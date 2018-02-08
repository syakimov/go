package main

import (
	"fmt"
)

func Log(s string) int {
	fmt.Println(s)
	return 42
}

func main() {
	Log("42")
}
