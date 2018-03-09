package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	// 2 1 4 0
}

// nums is a sorted array
func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < target && target <= nums[i] {
			return i
		}
	}

	return len(nums)
}
