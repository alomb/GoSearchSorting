package test

import (
	"fmt"
	"gosearchsorting/utils"
	"sort"
)

func Test(function func([]int, int) []int) {
	array := []int{4, 1, 2, 7, 5, 9, 8, 3, 10, 6, 3, 8}
	arraysorted := sort.IntSlice(array)
	result := function(array, len(array))

	fmt.Printf("result = %t (%v)", utils.Equals(result, arraysorted), result)
}
