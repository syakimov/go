// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 1}))
	fmt.Println(removeDuplicates([]int{1, 1, 3}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 3}))
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3}))
}

func removeDuplicatesSlower(nums []int) int {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == nums[i+1] {
			// this is very slow
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	sort.Ints(nums)
	return len(nums)
}

func removeDuplicates(nums []int) int {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == nums[i+1] {
			nums[i+1] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		}
	}

	sort.Ints(nums)
	return len(nums)
}
