package main

import (
	"fmt"
	"gosearchsorting/sorting/test"
	"sort"
)

// QuickSort in the classical recursive 2-way form
func QuickSort(array sort.Interface) sort.Interface {
	quicksort(array, 0, array.Len()-1)
	return array
}

// Returns the correct index of the last element of the subarray
func getRightIndex(array sort.Interface, min, max int) int {
	index := min

	// Find the index moving greater values to the right of the index and lower to the left
	for j := min; j < max; j++ {
		if array.Less(j, max) {
			array.Swap(j, index)
			index++
		}
	}

	// Finally put the value in the right position: this is the real ordering step
	array.Swap(index, max)
	return index
}

// The algorithm core
func quicksort(array sort.Interface, min, max int) {
	// If the array has length strictly greater than 1
	if min < max {
		// Compute the correct index
		index := getRightIndex(array, min, max)

		// From latin <3 "Divide et impera"
		quicksort(array, min, index-1)
		quicksort(array, index+1, max)
	}
}

func main() {
	fmt.Println("QuickSort")
	test.Test(QuickSort)
}
