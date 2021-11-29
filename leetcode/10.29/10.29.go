package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{2, 5, 6, 9, 10}
	fmt.Println(2 % 2)
	fmt.Println(findGCD(num))
}

func findGCD(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	max := 0
	for i := 1; i <= nums[0]; i++ {
		if nums[0]%i == 0 && nums[len(nums)-1]%i == 0 {
			max = i
		}
	}
	return max
}
