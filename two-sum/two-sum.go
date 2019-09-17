package twosum

import (
	"errors"
)

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	panic(errors.New("No two sum solution"))
}

func twoSumTwoPassHashTable(nums []int, target int) []int {
	numsMap := make(map[int]int)
	// In the first iteration, we add each element's value and its index to the table.
	for i, num := range nums {
		numsMap[num] = i
	}
	// In the second iteration we check if each element's complement (target−nums[i]) exists in the table.
	// Beware that the complement must not be nums[i] itself!
	var complement int
	for i, num := range nums {
		complement = target - num
		if (numsMap[complement] != 0) && (numsMap[complement] != i) {
			return []int{i, numsMap[complement]}
		}
	}
	panic(errors.New("No two sum solution"))
}
