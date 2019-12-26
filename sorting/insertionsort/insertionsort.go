package main

import (
	"fmt"
	"gosearchsorting/sorting/test"
	"gosearchsorting/utils"
)

func InsertionSort(array []int, length int) []int {
	// Outer loop to explore the unordered part of the array
	for i := 1; i < length; i++ {
		// Inner loop to insert an element in the ordered part of the array
		for j := i; j > 0; j-- {
			// Swaps until the i-th element is inserted in the right place
			if array[j] < array[j-1] {
				utils.Swap(array, j, j-1)
			} else {
				break
			}
		}
	}

	return array
}

func main() {
	fmt.Println("Insertion Sort")
	test.Test(InsertionSort)
}
