package main

import (
	"fmt"
	"gosearchsorting/sorting/test"
	"sort"
)

func SelectionSort(array sort.Interface) sort.Interface {
	// Outer loop (current element)
	for i := 0; i < array.Len()-1; i++ {
		minPosition := i
		// Inner loop (successors)
		for j := i; j < array.Len(); j++ {
			// If a successor is lower than the current element it will be swapped
			if array.Less(j, minPosition) {
				minPosition = j
			}
		}
		// Swap
		array.Swap(i, minPosition)
	}
	return array
}

func main() {
	fmt.Println("SelectionSort")
	test.Test(SelectionSort)
}
