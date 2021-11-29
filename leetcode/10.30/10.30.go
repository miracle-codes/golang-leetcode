package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(lenLongestFibSubseq(arr))
}
func lenLongestFibSubseq(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	//max:=0
	sum := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for h := j + 1; h < len(arr); h++ {
				//FibSubseq()
			}

		}
	}

	return sum
}

func FibSubseq(num []int, i int, j int, h int) (int, int) {
	if num[i]+num[j] == num[h] {
		return j, h
	}
	return i, j
}
