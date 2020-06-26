package provided

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	return fmt.Sprintf("%d -> %s", node.Val, node.Next)
}

func (curr *ListNode) Equal(other interface{}) bool {
	if other == nil {
		return false
	}
	node, ok := other.(*ListNode)
	if !ok {
		return false
	}
	if curr.Val != node.Val {
		return false
	}
	if curr.Next == nil || node.Next == nil {
		return curr.Next == nil && node.Next == nil
	}
	return curr.Next.Equal(node.Next)
}

func CreateListNode(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := new(ListNode)
	head.Val = vals[0]
	node := head
	for i := 1; i < len(vals); i++ {
		node.Next = new(ListNode)
		node = node.Next
		node.Val = vals[i]
	}
	return head
}
