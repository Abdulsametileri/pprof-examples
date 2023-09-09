package main

// The problem is here: https://leetcode.com/problems/two-sum/description/

/*
	Inlining refers to replacing a function call with the body of the function.
	Nowadays, inlining is done automatically by compilers.

	In order to get more accurate results, I disabled this compiler behaviour by providing
	`go:noinline`
*/

//go:noinline
func TwoSumWithBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return nil
}

//go:noinline
func TwoSumWithTwoPassHashTable(nums []int, target int) []int {
	indexByNums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		indexByNums[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if val, ok := indexByNums[complement]; ok && val != i {
			return []int{i, val}
		}
	}

	return nil
}

//go:noinline
func TwoSumOnePassHashTable(nums []int, target int) []int {
	numsByIndex := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if val, ok := numsByIndex[complement]; ok {
			return []int{val, i}
		}

		numsByIndex[nums[i]] = i
	}

	return nil
}

func main() {

}
