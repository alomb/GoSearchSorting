package utils

// Swap swaps two indexes in an array.
// Return true if possible, false otherwise.
func Swap(array []int, index1 int, index2 int) bool {
	if index1 >= 0 && index1 < len(array) && index2 >= 0 && index2 < len(array) && index1 != index2 {
		array[index1], array[index2] = array[index2], array[index1]
		return true
	} else {
		return false
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
