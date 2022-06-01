package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := addTree(arr, 0, len(arr)-1)
	//LevelTraversal(root)
	//RecursionPreorderTraversal(root)
	res := levelOrder(root)
	for _, v := range res {
		fmt.Println(v)
	}
}

type Node struct {
	Value       int
	Left, Right *Node
}

func addTree(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	root := &Node{Value: arr[mid]}
	root.Left = addTree(arr, start, mid-1)
	root.Right = addTree(arr, mid+1, end)
	return root

}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	q := []*Node{root}
	q[0].Value = 1
	//
	//nullRoot := &Node{Value: 0}
	//nullPointer :=[]*Node{nullRoot}
	for i := 0; len(q) > 0; i++ {
		result = append(result, []int{})
		p := []*Node{}
		for j := 0; j < len(q); j++ {
			result[i] = append(result[i], q[j].Value)
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
	}
	return result
}

func RecursionPreorderTraversal(node *Node) {
	if node == nil {
		return
	}
	RecursionPreorderTraversal(node.Left)
	fmt.Printf("%v  ", node.Value)
	RecursionPreorderTraversal(node.Right)

}

func LevelTraversal(rootNode *Node) {

	if rootNode == nil {
		return
	}

	//封装成slice
	var nodeSlice []*Node
	nodeSlice = append(nodeSlice, rootNode)

	//开始递归遍历
	RecursionTraversal(nodeSlice)
}

func RecursionTraversal(nodeSlice []*Node) {
	//如果当前层级slice为空，则结束遍历
	if len(nodeSlice) == 0 {
		return
	}

	//创建新的节点slice，存储下一层需要遍历的node
	var nextSlice []*Node

	//遍历当前nodeSlice
	for i := 0; i < len(nodeSlice); i++ {
		//取出要遍历的node
		node := nodeSlice[i]

		//输出当前node的值
		fmt.Printf("%v  ", node.Value)
		fmt.Println()
		//当前node左子节点append到下一层node slice中
		if node.Left != nil {
			nextSlice = append(nextSlice, node.Left)
		}

		//当前node右子节点append到下一层node slice中
		if node.Right != nil {
			nextSlice = append(nextSlice, node.Right)
		}
	}

	//递归遍历下一层的node slice
	RecursionTraversal(nextSlice)

}

//func levelOrder(root *Node) [][]int  {
//	res := [][]int{}
//	dfs(root,0)
//	return res
//}
//
//func dfs(root *Node,level int)  {
//	if root != nil {
//		if len(res) {
//
//		}
//	}
//
//}
