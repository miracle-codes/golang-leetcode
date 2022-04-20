package main

import "fmt"

type list struct {
	id   int
	val  int
	next *list
}

func insertList(list *list, arr []int) {
	temp := list
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for _, val := range arr {
		temp.next.val = val
		temp = temp.next
	}

}

func List(list *list) {
	temp := list
	if temp.val == 0 {
		fmt.Println("链表是空的")
		return
	}
	for {
		fmt.Printf("%d==>", temp.val)
		if temp.next == nil {
			break
		}
		temp = temp.next

	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	list := &list{}
	insertList(list, arr)
	List(list)
}
