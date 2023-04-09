package main

func romanToInt(s string) int {
	if len(s) > 15 || len(s) == 0 {
		return 0
	}
	romans := map[uint8]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	var result, latestVal, currentVal int
	for i := len(s) - 1; i >= 0; i-- {
		currentVal = romans[s[i]]
		if currentVal < latestVal {
			result -= currentVal
		} else {
			result += currentVal
		}
		latestVal = currentVal
	}

	return result
}
