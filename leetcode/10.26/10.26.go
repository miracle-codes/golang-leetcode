package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 6, 2, 7, 4}
	//quickSort(nums)
	//fmt.Println(nums)
	nums = append(nums[3:])
	fmt.Println(nums)
	fmt.Println(maxProductDifference(nums))
}

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}
func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := nums[0]
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	quickSort(nums[:i])
	quickSort(nums[i+1:])
}
