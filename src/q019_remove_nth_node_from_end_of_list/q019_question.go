// [Medium] Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
//
// Given a linked list, remove the n-th node from the end of list and return its head.
//
// Example:
//
//
// Given linked list: 1->2->3->4->5, and n = 2.
//
// After removing the second node from the end, the linked list becomes 1->2->3->5.
//
//
// Note:
//
// Given n will always be valid.
//
// Follow up:
//
// Could you do this in one pass?
package q019_remove_nth_node_from_end_of_list

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(*ListNode, int) *ListNode) {
	th := test.NewTestHelper(t)
	var res *ListNode

	defer th.NewTestCase()()
	res = method(CreateListNode(1, 2, 3, 4, 5), 2)
	th.AssertEqual(CreateListNode(1, 2, 3, 5), res)
}
