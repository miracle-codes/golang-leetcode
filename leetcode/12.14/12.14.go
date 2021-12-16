package main

import "fmt"

func main() {

	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt2(4))
}

func mySqrt(x int) int {
	left := 1
	for left*left <= x {
		left++
	}

	return left - 1
}

///二分法
func mySqrt2(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
