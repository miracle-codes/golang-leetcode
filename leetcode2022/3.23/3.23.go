package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6666, 1, 23, 3}
	target := 5
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	middle := len(nums) / 2
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[middle] > target {
			right = middle - 1
			middle = (left + right) / 2
		} else if nums[middle] < target {
			left = middle + 1
			middle = (left + right) / 2
		} else {
			return middle
		}
	}

	return -1
}
