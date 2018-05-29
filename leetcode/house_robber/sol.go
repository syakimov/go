package rob

// https://leetcode.com/problems/house-robber/description/

// Base Case: nums[0] = nums[0]
// nums[1] = max(nums[0], nums[1])
// nums[k] = max(k + nums[k-2], nums[k-1])
// Approach 1:- Construct dp table
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}

	return dp[len(dp)-1]
}

func robStollen(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}
	var a, b int

	for i := 0; i < len(nums); i++ {
		a, b = nums[i]+b, max(a, b)
	}

	return max(a, b)
}
