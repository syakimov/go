// https://leetcode.com/problems/two-sum/description/
package main

import (
	"fmt"
)

func main() {
	// expect [0 1]
	fmt.Println(twoSum([]int{2, 7, 15, 11}, 9))
	// expect [1 2]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

// Complexity: n
// Space: n
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)

	for k, v := range nums {
		another := dict[target-v]
		if dict[v] == 0 {
			dict[v] = k + 1
		}
		if another != 0 {
			return []int{another - 1, k}
		}
	}
	return []int{}
}

// Complexity: n * n
// Space: 0
func twoSumBruteforce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
