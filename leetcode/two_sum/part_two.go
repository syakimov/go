// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

package twosum

func twoSumBrute(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{0, 0}
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		switch {
		case sum < target:
			i++
		case sum > target:
			j--
		default:
			return []int{i + 1, j + 1}
		}
	}
	return []int{-1, -1}
}

func twoSumHash(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	for i, v := range numbers {
		if v1, found := m[target-v]; found {
			return []int{v1 + 1, i + 1}
		} else {
			m[v] = i
		}
	}
	return []int{}
}

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
