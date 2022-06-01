package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow(nums, k)
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	var temp = make([]int, 0)
	var res = make([]int, 0)

	for i := 0; i <= len(nums)-k; i++ {
		temp = nums[i:(i + k)]
		for _, v := range temp {
			fmt.Print(v)
		}
		fmt.Println()
		val := max(temp)
		res = append(res, val)
	}
	return res
}

func max(arr []int) int {
	res := math.MinInt32
	for _, v := range arr {
		res = maxVal(res, v)
	}
	return res
}

func maxVal(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
