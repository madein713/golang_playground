package main

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	idx := -1
	i := 0
	for i < len(haystack) {
		if haystack[i] == needle[0] && (i+len(needle) <= len(haystack)) {
			count := 0
			for j := 0; j < len(needle); j++ {
				if haystack[j+i] == needle[j] {
					count++
				}
			}
			if count == len(needle) {
				return i
			}

		}
		i += 1

	}
	return idx
	// Just use fucking slices instead of this shit
	//func strStr(haystack string, needle string) int {
	//    for i:= 0; i <= len(haystack) - len(needle); i++ {
	//        if haystack[i:len(needle)+i] == needle {
	//            return i
	//        }
	//    }
	//    return -1
	//}
}
