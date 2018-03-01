// https://leetcode.com/problems/longest-common-prefix/description/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"124", "12"}))
	// get runtime index out of range exception
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var i int
	for i = 0; i < len(strs[0]); i++ {
		c := strs[0][0]

		for k := 0; k < len(strs); k++ {
			if i > len(strs[k]) || c != strs[k][i] {
				return string(strs[0][:i+1])
			}
		}
	}

	return string(strs[0][:i])
}
