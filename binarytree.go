package datastructure

import (
	"fmt"
)
type BinaryTreeNode struct {
    Key int
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

func (n *BinaryTreeNode) TraverseInorder() {
	if n==nil {
		return
	}
	if n.Left != nil {
		n.Left.TraverseInorder()
	}
	fmt.Println(n.Key)
	if n.Right != nil {
	    n.Right.TraverseInorder()
	}
}

