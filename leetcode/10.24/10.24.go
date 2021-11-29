package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}
