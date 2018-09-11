package datastructure

import (
    "testing"
)

func TestInOrderTraversal(t *testing.T) {
    node1 := &BinaryTreeNode{Key:10}
	node2 := &BinaryTreeNode{Key:20}
	node3 := &BinaryTreeNode{Key:22}
	node4 := &BinaryTreeNode{Key:24}
	node5 := &BinaryTreeNode{Key:27}
	node6 := &BinaryTreeNode{Key:32}
	node7 := &BinaryTreeNode{Key:40}
	node8 := &BinaryTreeNode{Key:50}
	node9 := &BinaryTreeNode{Key:62}
	node5.Left = node3
	node5.Right = node8
	node3.Left = node2
	node3.Right = node4
	node2.Left = node1
	node8.Left = node6
	node8.Right = node9
	node6.Right = node7
	node5.TraverseInorder()
}
