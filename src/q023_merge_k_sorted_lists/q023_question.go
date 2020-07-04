// [Hard] Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists/
//
// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
// Example:
//
//
// Input:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// Output: 1->1->2->3->4->4->5->6
package q023_merge_k_sorted_lists

import (
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
	"github.com/sunxuia/leetcode-solution-go/src/util/test"
)

func doTest(t *testing.T, method func([]*ListNode) *ListNode) {
	th := test.NewTestHelper(t)
	var res *ListNode

	defer th.NewTestCase()()
	res = method([]*ListNode{CreateListNode(1, 4, 5), CreateListNode(1, 3, 4), CreateListNode(2, 6)})
	th.AssertEqual(CreateListNode(1, 1, 2, 3, 4, 4, 5, 6), res)

	defer th.NewTestCase()()
	res = method([]*ListNode{nil})
	th.AssertEqual(nil, res)
}
