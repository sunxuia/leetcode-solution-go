// [Easy] Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/
//
// Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.
//
// Example:
//
//
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
package q021_merge_two_sorted_lists

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func(*ListNode, *ListNode) *ListNode) {
	th := test.NewTestHelper(t)
	var res *ListNode

	defer th.NewTestCase()()
	res = method(CreateListNode(1, 2, 4), CreateListNode(1, 3, 4))
	th.AssertEqual(CreateListNode(1, 1, 2, 3, 4, 4), res)
}
