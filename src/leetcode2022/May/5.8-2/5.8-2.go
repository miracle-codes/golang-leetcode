package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 89}
	res := twoSum2(arr, 9)
	for _, v := range res {
		fmt.Println(v)
	}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for key, v := range nums {
		if preIndex, ok := m[target-v]; ok {
			return []int{preIndex, key}
		} else {
			m[v] = key
		}
	}
	return []int{}
}
