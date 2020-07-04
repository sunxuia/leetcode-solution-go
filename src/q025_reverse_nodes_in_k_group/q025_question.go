// [Hard] Reverse Nodes in k-Group
// https://leetcode.com/problems/reverse-nodes-in-k-group/
//
// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
//
//
//
// Example:
//
// Given this linked list: 1->2->3->4->5
//
// For k = 2, you should return: 2->1->4->3->5
//
// For k = 3, you should return: 3->2->1->4->5
//
// Note:
//
//
// 	Only constant extra memory is allowed.
// 	You may not alter the values in the list's nodes, only nodes itself may be changed.
package q025_reverse_nodes_in_k_group

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(*ListNode, int) *ListNode) {
	th := test.NewTestHelper(t)
	var res *ListNode

	res = method(CreateListNode(1, 2, 3, 4, 5), 2)
	th.AssertEqual(CreateListNode(2, 1, 4, 3, 5), res)

	res = method(CreateListNode(1, 2, 3, 4, 5), 3)
	th.AssertEqual(CreateListNode(3, 2, 1, 4, 5), res)
}
