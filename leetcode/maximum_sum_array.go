package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n, n)
	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = nums[i]

		if dp[i-1] > 0 {
			dp[i] = dp[i] + dp[i-1]
		}

		if dp[i] > max {
			max = dp[i]
		}
		fmt.Println(dp)
	}

	return max
}
