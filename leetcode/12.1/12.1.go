package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {

	fmt.Printf("a=%b,b=%b", a, b)
	fmt.Println()
	b = ^a
	fmt.Printf("a=%b,b=%b", a, b)

	fmt.Println()
	b = ^a
	fmt.Printf("a=%b,b=%b", a, b)
	//b ^=a
	//a ^=b
	return a, b
}

func main() {
	//fmt.Println(11<<1）
	//fmt.Println(16>>2)
	//11
	//1011
	//10000

	//fmt.Println(swap(-10,15))
	fmt.Println(divide(15, 2))
}
func divide(a int, b int) int {
	sign := 1
	if a^b < 0 {
		sign = -1
	}
	if a < 0 {
		a = ^a + 1
	}
	if b < 0 {
		b = ^b + 1
	}
	res := 0
	fmt.Printf("%b", a)
	for i := 31; i >= 0; i-- {
		fmt.Printf("%b    ", a>>i)
		fmt.Println(a, a>>i)
		//if a>>i >= b {
		//	if i == 31 && sign == 1 {
		//		return math.MaxInt32
		//	}
		//
		//	res += 1 << i
		//	a -= b << i
		//}
	}
	if sign == -1 {
		return ^res + 1
	}
	return res
}
