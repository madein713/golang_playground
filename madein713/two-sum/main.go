package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 15))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		for idx := 0; idx < len(nums); idx++ {
			var findSum int
			if a != nums[idx] {
				findSum = a + nums[idx]
			} else {
				continue
			}
			if findSum == target {
				return []int{i, idx}
			}
		}
	}
	return []int{}
}
