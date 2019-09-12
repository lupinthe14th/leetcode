package main

import "log"

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			log.Printf("i: %d, nums[i]: %v\n", i, nums[i])
			log.Printf("j: %d, nums[j]: %v\n", j, nums[j])
			sum := nums[i] + nums[j]
			log.Printf("sum: %d", sum)
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
