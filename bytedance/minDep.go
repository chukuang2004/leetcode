package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	NULL = 0xfffffffff
)

func array2tree(src []int) *TreeNode {

	size := len(src)
	if size == 0 {
		return nil
	}
	tmp := make([]*TreeNode, size)

	for i := 0; i < size; i++ {
		if src[i] != NULL {
			node := &TreeNode{}
			node.Val = src[i]
			tmp[i] = node
		} else {
			tmp[i] = nil
		}
	}

	for i := 0; i < size/2; i++ {
		if tmp[i] == nil {
			continue
		}

		tmp[i].Left = tmp[2*i+1]
		tmp[i].Right = tmp[2*i+2]
	}
	return tmp[0]
}

func traversalTree(root *TreeNode) {

	if root == nil {
		return
	}

	fmt.Println(root.Val)
	traversalTree(root.Left)
	traversalTree(root.Right)
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := 0
	right := 0
	child := 0

	if root.Left != nil {
		left = minDepth(root.Left)
	}

	if root.Right != nil {
		right = minDepth(root.Right)
	}

	if left > 0 && right > 0 {
		if left > right {
			child = right
		} else {
			child = left
		}
	} else {
		if left == 0 {
			child = right
		} else {
			child = left
		}
	}

	return 1 + child
}

func main() {

	// src := []int{1, 2, 3, 4, NULL, NULL, 15}
	// src := []int{1, NULL, 2}
	src := []int{3, 9, 20, NULL, NULL, 15, 7}
	root := array2tree(src)

	traversalTree(root)

	ret := minDepth(root)

	fmt.Println("-----------", ret)
}
