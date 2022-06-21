package tool_slice

// @Converter explicit:unique
func UniqueStringSlice(s1 []string) []string {
	uniqueMap := make(map[string]int)

	result := make([]string, 0, len(s1))
	for _, s := range s1 {
		if _, ok := uniqueMap[s]; ok {
			continue
		} else {
			result = append(result, s)
			uniqueMap[s] = 1
		}
	}

	return result
}

// @Converter explicit:unique
func UniqueIntSlice(s1 []int) []int {
	uniqueMap := make(map[int]int)

	result := make([]int, 0, len(s1))
	for _, s := range s1 {
		if _, ok := uniqueMap[s]; ok {
			continue
		} else {
			result = append(result, s)
			uniqueMap[s] = 1
		}
	}

	return result
}

// @Converter explicit:unique
func UniqueInt32Slice(s1 []int32) []int32 {
	uniqueMap := make(map[int32]int32)

	result := make([]int32, 0, len(s1))
	for _, s := range s1 {
		if _, ok := uniqueMap[s]; ok {
			continue
		} else {
			result = append(result, s)
			uniqueMap[s] = 1
		}
	}

	return result
}

// @Converter explicit:unique
func UniqueInt64Slice(s1 []int64) []int64 {
	uniqueMap := make(map[int64]int64)

	result := make([]int64, 0, len(s1))
	for _, s := range s1 {
		if _, ok := uniqueMap[s]; ok {
			continue
		} else {
			result = append(result, s)
			uniqueMap[s] = 1
		}
	}

	return result
}
