package main

import (
	"fmt"
	"sort"
)

// https://www.sohamkamani.com/blog/golang/arrays-vs-slices/
func main() {
	//arr := [4]int{3, 2, 5, 4}
	//longerArr := [5]int{5, 7, 1, 2, 0}

	//longerArr = arr
	//longerArr == arr

	type int4 struct {
		e0 int
		e1 int
		e2 int
		e3 int
	}

	type int5 struct {
		e0 int
		e1 int
		e2 int
		e3 int
		e5 int
	}

	//arr := int4{3, 2, 5, 4}
	//longerArr := int5{5, 7, 1, 2, 0}

	slice1 := []int{6, 1, 2}
	slice2 := []int{9, 3}

	// slices of any length can be assigned to other slice types
	slice1 = slice2
	fmt.Println("slice ==> ", slice1) // slice ==>  [9 3]

	nums := []int{8, 0}
	nums = append(nums, 8)
	fmt.Println("len() =>", len(nums), " cap() =>", cap(nums)) // len() => 3  cap() => 4

	// This creates a slice with length 0 and capacity 5
	arr := make([]int, 0, 5)
	arr = append(arr, 8)
	fmt.Println("arr ==> ", arr)
	arr[0] = 2
	fmt.Println("arr ==> ", arr)
	fmt.Println("len() =>", len(arr), " cap() =>", cap(arr)) // len() => 0  cap() => 5

	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)
}
