package main

import (
	"math"
)

func reverse(x int) int {
	var reversedNum int
	tmp := x
	multi := 1
	if tmp < 0 {
		multi = -1
		tmp = multi * tmp
	}
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		if multi*reversedNum >= math.MaxInt32 || multi*reversedNum <= math.MinInt32 {
			return 0
		}
		tmp = tmp / 10
	}
	return multi * reversedNum
}
