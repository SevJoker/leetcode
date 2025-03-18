package main

import (
	"fmt"
)


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//  二叉树中序遍历
// 中序    先左 再根 再右
// 前序    先根 再左 再右
// 后序	  先左 再右 再根

func inorderTraversal(root *TreeNode) []int {
	// if root == nil {
	// 	return []int{}
	// }
	// ret := inorderTraversalR(root.Left)
	// ret = append(ret, root.Val)
	// ret = append(ret,inorderTraversalR(root.Right)...)
	var ret = []int{}
	inorderTraversalR(root,&ret)
	return ret
}

func inorderTraversalR(root *TreeNode,ret *[]int) {
	if root == nil {
		return
	}
	inorderTraversalR(root.Left,ret)
	*ret = append(*ret, root.Val)
	inorderTraversalR(root.Right,ret)
	return
}

// 不用递归
func inorderTraversal1(root *TreeNode) (res []int) {
	stack := []*TreeNode{}

	for root!= nil||len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}

	return
}

func main() {

	n := &TreeNode{Val:1}
	n.Right = &TreeNode{Val:2}
	n.Right.Left = &TreeNode{Val:3}
	n.Right.Right = &TreeNode{Val:4}

	fmt.Println("---------------",inorderTraversal1(n))
}



// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/binary-tree-inorder-traversal/solutions/412886/er-cha-shu-de-zhong-xu-bian-li-by-leetcode-solutio/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func inorderTraversal2(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
