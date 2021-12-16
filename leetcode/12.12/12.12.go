package main

import "fmt"

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
	fmt.Printf("%v", arr)
}
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
