package main

import "fmt"

func twoSum(nums []int, target int) []int {
	size := len(nums)

	numbers := make(map[int]int)

	for i := 0; i < size; i++ {
		dif := target - nums[i]
		j, ok := numbers[dif]
		if ok {
			return []int{j, i}
		}
		numbers[nums[i]] = i
	}
	return nil
}

func twoSumBruteforce(nums []int, target int) []int {
	size := len(nums)
	for p1 := 0; p1 < size-1; p1++ {
		for p2 := 1; p2 < size; p2++ {
			if p1 == p2 {
				continue
			}
			if nums[p1]+nums[p2] == target {
				return []int{p1, p2}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("Input: nums = ", nums, " , target = ", target)
	result := twoSum(nums, target)
	fmt.Printf("Output: %v\n", result)
}
