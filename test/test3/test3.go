package main

import "fmt"

func main() {

	fmt.Println(test())
}

func test() int {
	res := 1
	defer func() {
		fmt.Println("start", res)
		res++
		fmt.Println("end", res)
	}()
	defer func() {
		fmt.Println("func1", res)
		res++
		fmt.Println("func2", res)
	}()
	defer func() {
		fmt.Println("func1", res)
		res++
		fmt.Println("func2", res)
	}()

	return 10
}
