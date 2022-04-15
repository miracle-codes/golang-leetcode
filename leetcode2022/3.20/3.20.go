package main

import "fmt"

func main() {
	num := []int{1, 5, 2, 3, 5, 6, 1, 23, 2, 5}
	fmt.Println(removeElement(num, 1))
}

func removeElement(nums []int, val int) int {

	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
			fmt.Println(nums)

		}
	}
	return left
}
