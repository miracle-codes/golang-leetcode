package main

import "fmt"

func main() {
	arr := []int{3, 2, 3, 4, 5, 1, 2, 3, 0, 123}

	fmt.Println(getLeastNumbers(arr, 1))
}

func getLeastNumbers(arr []int, k int) []int {
	arr2 := make([]int, len(arr))
	h := 0
	for i := 0; i < len(arr); h++ {
		min := arr[0]
		num := 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				num = j
			}
		}

		arr2[h] = arr[num]
		arr = append(arr[:num], arr[num+1:]...)
	}

	return arr2[:k]
}
