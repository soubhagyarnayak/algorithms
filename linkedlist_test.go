package datastructure

import (
	"testing"
)

func TestLinkedListTraversal(t *testing.T) {
	node1 := &LinkedListNode{Value: 10}
	node2 := &LinkedListNode{Value: 20}
	node3 := &LinkedListNode{Value: 22}
	node4 := &LinkedListNode{Value: 24}
	node5 := &LinkedListNode{Value: 27}
	node6 := &LinkedListNode{Value: 32}
	node7 := &LinkedListNode{Value: 40}
	node8 := &LinkedListNode{Value: 50}
	node9 := &LinkedListNode{Value: 62}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node9
	node1.PrintList()
}

func TestLinkedListCircle(t *testing.T) {
	node1 := &LinkedListNode{Value: 10}
	node2 := &LinkedListNode{Value: 20}
	node3 := &LinkedListNode{Value: 22}
	node4 := &LinkedListNode{Value: 24}
	node5 := &LinkedListNode{Value: 27}
	node6 := &LinkedListNode{Value: 32}
	node7 := &LinkedListNode{Value: 40}
	node8 := &LinkedListNode{Value: 50}
	node9 := &LinkedListNode{Value: 62}
	node1.Next = node2
	if got := node1.HasCircle(); got != false {
		t.Errorf("Expected no circle, but got circle")
	}
	if got := node3.HasCircle(); got != false {
		t.Errorf("Expected no circle, but got circle")
	}
	node4.Next = node5
	node5.Next = node6
	node6.Next = node9
	node9.Next = node7
	node7.Next = node8
	node8.Next = node5
	if got := node4.HasCircle(); got != true {
		t.Errorf("Expected circle, but got no circle")
	}
}
