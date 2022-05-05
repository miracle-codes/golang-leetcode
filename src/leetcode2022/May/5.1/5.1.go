package main

import (
	"leetcode2022/func/makeList"
	"leetcode2022/func/showList"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := makeList.MakeListNode(arr)
	showList.ShowList(list)

}
