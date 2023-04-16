package main

func removeDuplicates(nums []int) int {
	//if len(nums) == 1 {
	//	return 1
	//}
	//
	//for i := 0; i < len(nums); i++ {
	//	//mark duplicates
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i] == nums[j] {
	//			nums[j] = 500
	//		}
	//	}
	//	// insert sort
	//	for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
	//		nums[j], nums[j-1] = nums[j-1], nums[j]
	//	}
	//}
	//num := 0
	//
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] != 500 {
	//		num++
	//		continue
	//	}
	//	break
	//}
	//return num
	i := 0
	for j := range nums {
		if nums[i] != nums[j] {
			i += 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}
