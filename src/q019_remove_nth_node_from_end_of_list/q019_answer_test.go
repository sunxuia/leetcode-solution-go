package q019_remove_nth_node_from_end_of_list

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
)

func TestRemoveNthFromEnd(t *testing.T) {
	doTest(t, removeNthFromEnd)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ahead := head
	// (题目中保证了n 肯定是有效的)
	for n > 0 {
		ahead = ahead.Next
		n--
	}

	// 删掉开头
	if ahead == nil {
		return head.Next
	}

	// 删掉中间或末尾的一个值
	node := head
	for ahead.Next != nil {
		ahead = ahead.Next
		node = node.Next
	}
	node.Next = node.Next.Next
	return head
}
