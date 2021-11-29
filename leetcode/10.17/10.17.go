package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 2, 1, 5, 3, 5, 6}
	//num = append(num[:1],num[2:]...)
	//fmt.Println(num)
	fmt.Println(singleNonDuplicate(num))
}

func singleNonDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		judge := false
		value := nums[i]

		for j := i + 1; j < len(nums); j++ {
			if nums[j] == value {
				judge = true
				nums = append(nums[:j], nums[j+1:]...)
				nums = append(nums[:i], nums[i+1:]...)
				i--
				break
			}
		}

		if judge == false {
			return nums[i]
		}
	}

	return 1
}
