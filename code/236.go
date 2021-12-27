package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //  递归解法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}


	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)

	if left != nil  && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil

}
func main() {

}