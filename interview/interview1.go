package main

import "fmt"

func main() {
	defer_call()
}
func defer_call() {
	defer func() {
		fmt.Println("1111")
	}()
	defer func() {
		fmt.Println("2222")

	}()
	defer func() {
		fmt.Println("3333")
	}()

	panic("lalalala")
}
