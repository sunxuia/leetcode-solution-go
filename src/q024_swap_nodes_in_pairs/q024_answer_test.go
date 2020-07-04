package q024_swap_nodes_in_pairs

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
)

func TestSwapPairs(t *testing.T) {
	doTest(t, swapPairs)
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	curr, next := dummy, dummy.Next
	for next != nil && next.Next != nil {
		curr.Next = next.Next
		next.Next, next.Next.Next = next.Next.Next, next
		curr, next = next, next.Next
	}
	return dummy.Next
}
