// https://leetcode.com/problems/longest-common-prefix/description/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"124", "12"}))
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
	fmt.Println(longestCommonPrefix([]string{"a", "a"}))
	fmt.Println(longestCommonPrefix([]string{"a", "b"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
