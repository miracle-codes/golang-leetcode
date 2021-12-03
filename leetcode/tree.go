package main

import "fmt"

//Cat 猫猫结构体
type Cat struct {
	No    int
	Name  string
	Left  *Cat
	Right *Cat
}

//EachByLast 前序遍历,先输出root节点,在输出左侧子树,在输出右侧子树
func EachByLast(catNode *Cat) {
	if catNode != nil {
		fmt.Printf("%v %v\n", catNode.No, catNode.Name)
		EachByLast(catNode.Left)
		EachByLast(catNode.Right)
	}
}

//EachByInfixOrder 中序遍历,先输出root左侧子树,在输出root节点,在输出root右侧子树
func EachByInfixOrder(catNode *Cat) {
	if catNode != nil {

		EachByInfixOrder(catNode.Left)
		fmt.Printf("%v %v\n", catNode.No, catNode.Name)
		EachByInfixOrder(catNode.Right)
	}
}

//EachByPostOrder 后序遍历,先输出root右侧子树,在输出root节点,在输出root左侧子树
func EachByPostOrder(catNode *Cat) {
	if catNode != nil {

		EachByPostOrder(catNode.Left)
		fmt.Printf("%v %v\n", catNode.No, catNode.Name)
		EachByPostOrder(catNode.Right)

	}
}

//4 三花猫
//3 小黑猫
//1 汤姆猫
//2 小白猫
//6 奶牛猫
//5 大橘猫

func main() {
	//构建一个二叉树
	root := &Cat{
		No:   1,
		Name: "汤姆猫",
	}
	left1 := &Cat{
		No:   2,
		Name: "小白猫",
	}
	left2 := &Cat{
		No:   5,
		Name: "大橘猫",
	}

	right1 := &Cat{
		No:   3,
		Name: "小黑猫",
	}
	right2 := &Cat{
		No:   4,
		Name: "三花猫",
	}
	right3 := &Cat{
		No:   6,
		Name: "奶牛猫",
	}
	root.Left = left1
	left1.Left = left2
	left2.Right = right3
	root.Right = right1
	right1.Right = right2

	EachByPostOrder(root)
}
