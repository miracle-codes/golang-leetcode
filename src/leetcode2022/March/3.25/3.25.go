package main

func main() {

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
