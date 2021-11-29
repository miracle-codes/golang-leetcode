package main

import (
	"strconv"
)

func main() {
	println(monotoneIncreasingDigits(987564))
}
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	num := len(s)
	if num == 1 {
		return n
	}
	arr := make([]int, num)
	for i := num - 1; i >= 0; i-- {
		arr[i], _ = strconv.Atoi(s[i : i+1])
	}

	color := false
	for color == false {
		color = true
		for j := num - 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				color = false
				arr[j-1] = arr[j-1] - 1
				for h := j; h < num; h++ {
					arr[h] = 9
				}
			}

		}
	}
	h := ""

	for key := range arr {
		h += strconv.Itoa(arr[key])
	}
	result, _ := strconv.Atoi(h)

	return result
}
