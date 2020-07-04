// [Medium] Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/
//
// Given a linked list, swap every two adjacent nodes and return its head.
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
//
//
// Example:
//
//
// Given 1->2->3->4, you should return the list as 2->1->4->3.
package q024_swap_nodes_in_pairs

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(*ListNode) *ListNode) {
	th := test.NewTestHelper(t)
	var res *ListNode

	defer th.NewTestCase()()
	res = method(CreateListNode(1, 2, 3, 4))
	th.AssertEqual(CreateListNode(2, 1, 4, 3), res)
}
