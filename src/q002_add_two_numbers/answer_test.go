package q002_add_two_numbers

import (
	"testing"

	. "../util/provided"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	node := dummy
	carr := 0
	n1, n2 := l1, l2
	for n1 != nil || n2 != nil {
		var v1, v2 int
		if n1 != nil {
			v1 = n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			v2 = n2.Val
			n2 = n2.Next
		}
		node.Next = &ListNode{Val: (carr + v1 + v2) % 10}
		carr = (carr + v1 + v2) / 10
		node = node.Next
	}
	if carr > 0 {
		node.Next = &ListNode{Val: carr}
	}
	return dummy.Next
}

func TestAddTwoNumbers(t *testing.T) {
	runner.Test(t, addTwoNumbers)
}
