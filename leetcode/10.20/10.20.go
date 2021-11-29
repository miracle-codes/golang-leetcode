package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}
	//fmt.Printf("$v", runningSum(num))
	fmt.Println(runningSum(num))
}

func runningSum(nums []int) []int {

	nums2 := nums
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums2[i] = nums[i]
		} else {
			nums2[i] = nums[i] + nums2[i-1]
		}
	}
	return nums2
}
