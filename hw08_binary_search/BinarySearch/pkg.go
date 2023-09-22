package binarysearch

func BinarySearch(array []int, key int) (bool, int) {
	var mid, exp int
	min := 0
	high := len(array) - 1
	for min <= high {
		mid = (min + high) / 2
		exp = mid
		if array[exp] == key {
			return true, exp
		}
		if array[exp] > key {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return false, -1
}
