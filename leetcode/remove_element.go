package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		}
	}

	fmt.Println(nums)
	return len(nums)
}
