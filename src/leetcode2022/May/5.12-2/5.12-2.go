package main

import "fmt"

func main() {
	fmt.Println(reverseStr("abcd", 4))
}
func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		fmt.Println(i)
		if i+k <= length {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func reverseStr2(s string, k int) string {
	arr := make([]byte, len(s))
	for k, v := range s {
		arr[k] = byte(v)
	}
	for i := 1; i <= len(s); i++ {
		if i%(2*k) == 0 {
			n := i / (2 * k)
			left := i - 2*k*n
			right := i - k*n - 1
			fmt.Printf("left = %d, right = %d", left, right)
			fmt.Println()
			for left < right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		if i == len(s) && i%(2*k) > k {
			n := i / (2 * k)
			left := 2 * k * n
			right := 2*k*n + k - 1
			fmt.Printf("left = %d, right = %d", left, right)
			fmt.Println()
			for left < right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		if i == len(s) && i == k {
			left := 0
			right := k - 1
			fmt.Printf("left = %d, right = %d", left, right)
			fmt.Println()
			for left < right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
	}
	var res = string(arr)

	return res

}
