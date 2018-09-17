package datastructure

import (
	"fmt"
)

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

func (list *LinkedListNode) PrintList() {
	var current = list
	for current != nil {
		fmt.Printf("%d->", current.Value)
		current = current.Next
	}
}

func (list *LinkedListNode) HasCircle() bool {
	var step1 = list
	var step2 = list
	for step1 != nil && step2 != nil && step2.Next != nil {
		step1 = step1.Next
		step2 = step2.Next.Next
		if step1 == step2 {
			return true
		}
	}
	return false
}
