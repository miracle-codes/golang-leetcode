package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	arr2 := []int{1, 2, 3, 7, 111, 1234, 1, 1, 1, 1}

	fmt.Println("arr1数组")
	for _, val := range arr1 {
		fmt.Printf("%d  ", val)
	}
	fmt.Println()

	fmt.Println("arr2数组")
	for _, val := range arr2 {
		fmt.Printf("%d  ", val)
	}
	fmt.Println()

	fmt.Println("重叠的数组")

	res := intersection(arr1, arr2)
	for _, val := range res {
		fmt.Printf("%d  ", val)
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}
	var res []int
	// 利用count>0，实现重复值只拿一次放入返回结果中
	for _, v := range nums2 {
		if count, ok := m[v]; ok && count > 0 {
			res = append(res, v)
			fmt.Printf("count = %d,ok = %v", count, ok)
			fmt.Println()
			m[v]--
		}
	}
	return res

}
