package main

// This problem is practically Fibonacci.
// You have base cases for n <= 2
// n = n-2 + n-1
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	oneStepBefore := 2
	twoStepsBefore := 1
	current := 0

	for i := 2; i < n; i++ {
		current = oneStepBefore + twoStepsBefore
		twoStepsBefore = oneStepBefore
		oneStepBefore = current
	}

	return current
}
