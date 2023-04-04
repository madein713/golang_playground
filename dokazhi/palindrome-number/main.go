package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	stri := strconv.Itoa(x)
	equal := true
	for i := 0; i < (len(stri)/2)+1; i++ {
		fmt.Printf("%c: %d ==== %c: %d\n", stri[i], i, stri[len(stri)-1-i], len(stri)-1-i)
		if string(stri[i]) != string(stri[len(stri)-1-i]) {
			equal = false
			break
		}
	}
	return equal
}

func isPalindromeRight(x int) bool {
	var reversedNum int
	tmp := x
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}
	return x == reversedNum
}

func main() {
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))

	fmt.Println(isPalindromeRight(12321))
	fmt.Println(isPalindromeRight(-121))
	fmt.Println(isPalindromeRight(10))
}
