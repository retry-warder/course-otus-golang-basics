package binarysearch

func BinarySearch(bsa []int, searchval int) (bool, int) {
	var mid int
	min := 0
	high := len(bsa) - 1
	if searchval < 0 || searchval > high {
		return false, -1
	}
	for min <= high {
		mid = (min + high) / 2
		if mid == searchval {
			return true, mid
		}
		if mid > searchval {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return false, -1
}
