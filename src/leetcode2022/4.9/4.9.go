package main

import "fmt"

func main() {

	fmt.Println(generateMatrix(4))
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	up := 0
	down := n - 1
	left := 0
	right := n - 1
	index := 1

	for index <= n*n {

		for i := left; i <= right; i++ {
			res[up][i] = index
			index++
		}
		up++
		for i := up; i <= down; i++ {
			res[i][right] = index
			index++
		}
		right--
		for i := right; i >= left; i-- {
			res[down][i] = index
			index++
		}
		down--
		for i := down; i >= up; i-- {
			res[i][left] = index
			index++
		}
		left++
	}
	//for i:= range res{
	//	for j:=range res[i]{
	//		if res[i][j]<10 {
	//			fmt.Printf("%v  ",res[i][j])
	//		}else {
	//			fmt.Printf("%v ",res[i][j])
	//		}
	//	}
	//	fmt.Println()
	//}
	return res

}

//1  2  3  4
//12 13 14 5
//11 16 15 6
//10  9  8 7
