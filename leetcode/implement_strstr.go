// https://leetcode.com/problems/implement-strstr/description/
package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}

func strStr(haystack, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
