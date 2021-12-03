package main

import "fmt"

type TreeNode struct {
	val   int
	index int
	left  *TreeNode
	right *TreeNode
}

func main() {
	root := createTreeNode([]int{8, 4, 9, 3, 5, -1, 10})
	fmt.Println(root.val)
	fmt.Println(root.left.val, root.right.val)

	fmt.Println(root.left.left.val, root.left.right.val, root.right.right.val)
}

// BFS的数组翻译为二叉树
func createTreeNode(nums []int) *TreeNode {
	var root *TreeNode
	var queue []*TreeNode
	for i := 0; i < len(nums); i += 2 {
		if i == 0 {
			root = &TreeNode{val: nums[i], index: i}
			queue = append(queue, root)
		}
		parentNode := queue[0]
		queue = queue[1:]

		// 添加左节点
		if i+1 < len(nums) && nums[i+1] != -1 {
			parentNode.left = &TreeNode{val: nums[i+1], index: i + 1}
			queue = append(queue, parentNode.left)
		}
		//if  nums[i+1] == -1 {
		//	parentNode.left = &TreeNode{val: 0, index: i + 1}
		//	queue = append(queue, parentNode.left)
		//}

		// 添加右节点
		if i+2 < len(nums) && nums[i+2] != -1 {
			parentNode.right = &TreeNode{val: nums[i+2], index: i + 2}
			queue = append(queue, parentNode.right)
		}
	}
	return root
}
