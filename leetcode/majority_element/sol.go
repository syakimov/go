package majorityel

// https://leetcode.com/problems/majority-element/description/
func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++

		if m[n] > len(nums)/2 {
			return n
		}
	}

	return 42
}

// https://leetcode.com/problems/factorial-trailing-zeroes/description/
func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	return n/5 + trailingZeroes(n/5)
}

func trailingZeroesIterative(n int) int {
	res := 0
	for n > 0 {
		res, n = res+n/5, n/5
	}
	return res
}
