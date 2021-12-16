package main

import "fmt"

func main() {
	arr := []int{1, 2, 0, 0, 9, 10}
	moveZeroes(arr)
	fmt.Printf("%v", arr)
}

func moveZeroes(nums []int) {
	left := 0
	for _, v := range nums {
		if v != 0 {
			nums[left] = v
			left++
		}
	}
	for i := left; i <= len(nums)-1; i++ {
		nums[i] = 0
	}

}
