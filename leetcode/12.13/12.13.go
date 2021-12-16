package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
}

func isPerfectSquare(num int) bool {
	l, r := 0, num
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= num {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if ans*ans == num {
		return true
	} else {
		return false
	}

}
