package q021_merge_two_sorted_lists

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
)

func TestMergeTwoLists(t *testing.T) {
	doTest(t, mergeTwoLists)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 == nil {
		curr.Next = l2
	} else {
		curr.Next = l1
	}
	return dummy.Next
}
