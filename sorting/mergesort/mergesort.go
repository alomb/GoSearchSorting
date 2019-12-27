package main

import (
	"fmt"
	"gosearchsorting/sorting/test"
	"sort"
)

// Classical recursive MergeSort.
// Merge is in place.
func MergeSort(array sort.Interface) sort.Interface {
	mergesort(array, 0, array.Len()-1)
	return array
}

// Merge and order the two subarray [min...medium] and [medium+1...max]
func merge(array sort.Interface, start, medium, end int) {
	start2 := medium + 1

	// If the direct merge is already sorted
	if array.Less(medium, start2) {
		return
	}

	// Two pointers to maintain start
	// of both arrays to merge
	for start <= medium && start2 <= end {

		// If element on the left is less than the other on the right keep going
		if array.Less(start, start2) {
			start++
		} else {
			index := start2

			// Otherwise shift all elements from the two starts swapping them
			for index != start {
				array.Swap(index, index-1)
				index--
			}

			start++
			medium++
			start2++
		}
	}
}

// Algorithm core
func mergesort(array sort.Interface, start, end int) {
	if start < end {
		medium := start + (int)(end-start)/2
		mergesort(array, start, medium)
		mergesort(array, medium+1, end)

		merge(array, start, medium, end)
	}
}

func main() {
	fmt.Println("MergeSort")
	test.Test(MergeSort)
}
