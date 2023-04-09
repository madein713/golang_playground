package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// this is O(n2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumOn(nums []int, target int) []int {
	s := make(map[int]int)
	for id, num := range nums {
		if pos, ok := s[target-num]; ok {
			// если в сете есть цифра которая получается путем вычитания текущей цифры из искомой то мы знаем 2 индекса
			// 1 - это текущий индекс, 2 - тот который есть в сете путем вычитания
			fmt.Println(pos, ok)
			return []int{pos, id}
		}
		s[num] = id
		fmt.Println(s)
	}
	return []int{}
}
