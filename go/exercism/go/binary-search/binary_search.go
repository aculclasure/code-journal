package binarysearch

// SearchInts accepts a list of numbers and a key, executes
// a binary search on the list for the key and returns the
// index position of the key in the list if it's found. If
// the key is not located, then -1 is returned.
func SearchInts(values []int, key int) int {
	if len(values) == 0 {
		return -1
	}

	head, tail := 0, len(values)-1
	for {
		if values[head] == key {
			return head
		}

		if head >= tail || tail < 0 {
			return -1
		}

		mid := (tail-head)/2 + head
		switch {
		case key == values[mid]:
			return mid
		case key > values[mid]:
			head = mid + 1
		default:
			tail = mid - 1
		}
	}
}
