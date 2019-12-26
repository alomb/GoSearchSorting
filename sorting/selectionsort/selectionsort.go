package main

import (
	"fmt"
	"gosearchsorting/sorting/test"
	"gosearchsorting/utils"
)

func SelectionSort(array []int, length int) []int {
	// Outer loop (current element)
	for i := 0; i < length-1; i++ {
		minPosition := i
		// Inner loop (successors)
		for j := i; j < length; j++ {
			// If a successor is lower than the current element it will be swapped
			if array[j] < array[minPosition] {
				minPosition = j
			}
		}
		// Swap
		utils.Swap(array, i, minPosition)
	}
	return array
}

func main() {
	fmt.Println("Selection Sort")
	test.Test(SelectionSort)
}
