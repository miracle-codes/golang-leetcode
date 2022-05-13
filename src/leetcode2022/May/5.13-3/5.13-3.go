package main

import "fmt"

func main() {
	s := "asdqwezxc"
	fmt.Println(reverseLeftWords(s, 2))
}

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	rever(b, 0, n-1)
	rever(b, n, len(b)-1)
	rever(b, 0, len(b)-1)
	return string(b)
}

func rever(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
