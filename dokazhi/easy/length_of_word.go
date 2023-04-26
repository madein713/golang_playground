package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	strs := strings.Split(s, " ")
	gotLast := true
	lastIndex := len(strs) - 1
	for gotLast {
		//gotLast = true
		if strs[lastIndex] == "" {
			lastIndex -= 1
			continue
		}
		gotLast = false
		return len(strs[lastIndex])
	}
	return 0
}
