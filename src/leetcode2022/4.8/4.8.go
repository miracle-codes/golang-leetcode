package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	top := 0
	left := 0
	right := n - 1
	bottom := n - 1
	val := 1
	for val <= n*n {
		for i := left; i <= right; i++ {
			arr[top][i] = val
			val++
		}
		top++
		for i := top; i <= bottom; i++ {
			arr[i][right] = val
			val++
		}
		right--
		for i := right; i >= left; i-- {
			arr[bottom][i] = val
			val++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			arr[i][left] = val
			val++
		}
		left++
	}

	return arr
}
