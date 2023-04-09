package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(122))
	fmt.Println(isPalindrome(1212))
	fmt.Println(isPalindrome(12121))
}
func isPalindrome(num int) bool {
	input_num := num
	var remainder int
	res := 0
	for num > 0 {
		remainder = num % 10
		res = (res * 10) + remainder
		num = num / 10
	}
	return input_num == res
}
