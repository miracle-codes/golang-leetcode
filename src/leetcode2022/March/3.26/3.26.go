package main

import (
	"fmt"
	"sort"
)

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		fmt.Println(key)
		fmt.Println(arr)
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func square(arr []int) []int {
	for key, val := range arr {
		arr[key] = val * val
	}

	sort.Ints(arr)
	return arr
}

func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	fmt.Println(square(arr))
	//quickSort(arr, 0, len(arr)-1)
	//fmt.Println(arr)
}
