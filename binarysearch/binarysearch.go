package binarysearch

func Search(q int, data []int) bool {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + (high-low)/2
		if q < data[mid] {
			high = mid - 1
			continue
		}
		if q > data[mid] {
			low = mid + 1
			continue
		}
		return true
	}
	return false
}
