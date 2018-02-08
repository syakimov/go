// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import "fmt"
import "time"

func main() {
	start := time.Now()

	c := make(chan int)

	go func(c chan int) {
		z := 0
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			z += i
		}
		c <- z
	}(c) // Note the parentheses - must call the function.

	go func(c chan int) {
		z := 0
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			z += i
		}
		c <- z
	}(c) // Note the parentheses - must call the function.

	time.Sleep(1100 * time.Millisecond)

	<-c
	fmt.Println(time.Now().Sub(start))
}
