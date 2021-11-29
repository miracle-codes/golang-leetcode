package main

import "fmt"

func main() {
	num := []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(num, 1))
}
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 1 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if j-i <= k {
					return true
				}
			}
		}
	}
	return false

}
