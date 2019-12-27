package test

import (
	"fmt"
	"sort"
)

// Base test
func Test(function func(sort.Interface) sort.Interface) {
	test1 := sort.IntSlice{4, 1, 2, 7, 5, 9, 8, 3, 10, 6, 3, 8}
	result := function(test1)

	fmt.Printf("result = %t (%v)", sort.IsSorted(result), result)
}
