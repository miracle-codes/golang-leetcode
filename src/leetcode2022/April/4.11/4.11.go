package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, arr))
}

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := len(nums)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= s {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//func minSubArrayLen(target int, nums []int) int {
//	sum := 0
//	long :=0
//	res := len(nums)
//	for _,val:=range nums{
//		sum+=val
//		if sum>target {
//			if res>long {
//				res = long
//				sum = 0
//			}else {
//				sum =0
//				long =0
//			}
//
//		}
//		long++
//	}
//	if res==len(nums) {
//		return 0
//	}
//
//	return res
//}
