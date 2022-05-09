package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	arr2 := []int{1, 2, 3, 7, 111, 1234, 1, 1, 1, 1}

	res := intersection(arr1, arr2)
	for _, val := range res {
		fmt.Println(val)
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var res []int
	for _, val := range nums1 {
		m[val] = 1
	}

	for _, val := range nums2 {
		if m[val] > 0 {
			res = append(res, val)
			m[val]--
		}
	}
	return res

}
