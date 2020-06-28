// Package q002_add_two_numbers [Medium] Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/
//
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
//
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
package q002_add_two_numbers

import (
	"testing"

	"github.com/sunxuia/leetcode-solution-go/src/util/test"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
)

func doTest(t *testing.T, method func(*ListNode, *ListNode) *ListNode) {
	ah := test.NewAssertHelper()
	var res *ListNode

	res = method(CreateListNode(2, 4, 3), CreateListNode(5, 6, 4))
	ah.Assert(CreateListNode(7, 0, 8), res)

	res = method(CreateListNode(5), CreateListNode(5))
	ah.Assert(CreateListNode(0, 1), res)

	res = method(CreateListNode(1, 8), CreateListNode(0))
	ah.Assert(CreateListNode(1, 8), res)
}
