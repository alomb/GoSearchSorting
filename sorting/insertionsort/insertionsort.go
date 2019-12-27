package main

import (
	"fmt"
	"gosearchsorting/sorting/test"
	"sort"
)

func InsertionSort(array sort.Interface) sort.Interface {
	// Outer loop to explore the unordered part of the array
	for i := 1; i < array.Len(); i++ {
		// Inner loop to insert an element in the ordered part of the array
		for j := i; j > 0; j-- {
			// Swaps until the i-th element is inserted in the right place
			if array.Less(j, j-1) {
				array.Swap(j, j-1)
			} else {
				break
			}
		}
	}

	return array
}

func main() {
	fmt.Println("InsertionSort")
	test.Test(InsertionSort)
}
