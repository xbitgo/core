package tool_slice

func ContainsInt64(slice []int64, needle int64) bool {
	for _, item := range slice {
		if item == needle {
			return true
		}
	}
	return false
}

func ContainsInt32(slice []int32, needle int32) bool {
	for _, item := range slice {
		if item == needle {
			return true
		}
	}
	return false
}

func ContainsInt(slice []int, needle int) bool {
	for _, item := range slice {
		if item == needle {
			return true
		}
	}
	return false
}

func ContainsString(slice []string, needle string) bool {
	for _, item := range slice {
		if item == needle {
			return true
		}
	}
	return false
}
