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
	// Just use fucking slices instead of this shit this is from solutions
	//func strStr(haystack string, needle string) int {
	//    for i:= 0; i <= len(haystack) - len(needle); i++ {
	//        if haystack[i:len(needle)+i] == needle {
	//            return i
	//        }
	//    }
	//    return -1
	//}

	//this if from ChatGPT
	//func strStr(haystack string, needle string) int {
	//    if len(haystack) < len(needle) {
	//        return -1
	//    }
	//    if needle == "" {
	//        return 0
	//    }
	//    // build prefix function
	//    prefix := make([]int, len(needle))
	//    j := 0
	//    for i := 1; i < len(needle); i++ {
	//        for j > 0 && needle[i] != needle[j] {
	//            j = prefix[j-1]
	//        }
	//        if needle[i] == needle[j] {
	//            j++
	//        }
	//        prefix[i] = j
	//    }
	//    // match strings
	//    j = 0
	//    for i := 0; i < len(haystack); i++ {
	//        for j > 0 && haystack[i] != needle[j] {
	//            j = prefix[j-1]
	//        }
	//        if haystack[i] == needle[j] {
	//            j++
	//        }
	//        if j == len(needle) {
	//            return i - j + 1
	//        }
	//    }
	//    return -1
	//}
}
