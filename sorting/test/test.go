package test

import (
	"fmt"
)

func Test(function func([]int, int) []int) {
	array := []int{4, 1, 2, 7, 5, 9, 8, 3, 10, 6}

	result := function(array, len(array))
	fmt.Printf("%v", result)
}
