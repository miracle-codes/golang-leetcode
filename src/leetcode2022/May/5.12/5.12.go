package main

import "fmt"

func main() {
	arr := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(arr)
	fmt.Println(string(arr))
}

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

}
