package main

import "fmt"

func main() {
	fmt.Println(function4(2, 10))
}

///求x的n次方
//时间复杂度o(logn)
func function4(x int, n int) int {
	if n == 0 {
		return 1
	}

	t := function4(x, n/2)
	if n%2 == 1 {
		return t * t * x
	}
	return t * t
}
