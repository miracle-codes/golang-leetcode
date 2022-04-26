package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 5, 56, 6, 1}
	fmt.Println(minSubArrayLen(7, nums))
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				res = min(res, j-i+1)
			}

		}

	}
	if res == math.MaxInt32 {
		return 0
	}
	return res

}
func min(i int, j int) int {
	if i > j {
		return j
	}
	return i

}
