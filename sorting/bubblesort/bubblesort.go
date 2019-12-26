package main

import (
	"fmt"
	"gosearchsorting/sorting/test"
)

func BubbleSort(array []int, length int) []int {
	// Keep in memory the number of iterations to avoid useless checks
	iteration := 0
	for {
		var swapped bool
		// Iterate to push greater values to the end of the array
		for i := 0; i < length-1-iteration; i++ {
			// Swap if the previous element is greater than the following
			if array[i] > array[i+1] {
				tmp := array[i]
				array[i] = array[i+1]
				array[i+1] = tmp
				swapped = true
			}
		}

		// If no swap occured the array is ordered
		if !swapped {
			break
		}

		iteration++
	}

	return array
}

func main() {
	fmt.Println("Bubble Sort")
	test.Test(BubbleSort)
}
