package q025_reverse_nodes_in_k_group

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
)

func TestReverseKGroup(t *testing.T) {
	doTest(t, reverseKGroup)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	list := make([]*ListNode, 0, k)
	dummy := &ListNode{Val: 0, Next: head}
	prev, node := dummy, head
	for node != nil {
		for i := 0; i < k && node != nil; i++ {
			list = append(list, node)
			node = node.Next
		}
		if len(list) == k {
			for j := k - 1; j > 0; j-- {
				list[j].Next = list[j-1]
			}
			list[0].Next = node
			prev.Next = list[k-1]
			prev = list[0]
			list = list[:0]
		}
	}
	return dummy.Next
}

func TestReverseKGroup2(t *testing.T) {
	doTest(t, reverseKGroup2)
}

// 不使用数组缓存的方式
func reverseKGroup2(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev, node := dummy, head
	for node != nil {
		i := 0
		for ; i < k && node != nil; i++ {
			node = node.Next
		}
		if i == k {
			tail, curr := prev.Next, prev.Next
			prev.Next = node
			for curr != node {
				prev.Next, curr.Next, curr = curr, prev.Next, curr.Next
			}
			prev = tail
		}
	}
	return dummy.Next
}
