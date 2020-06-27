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
	. "../util/provided"
	"../util/test"
)

var runner = test.NewMethodTestRunner().AddTestCase(
	test.NewMethodTestCase(CreateListNode(2, 4, 3), CreateListNode(5, 6, 4)).ExpectReturn(CreateListNode(7, 0, 8))).AddTestCase(
	test.NewMethodTestCase(CreateListNode(5), CreateListNode(5)).ExpectReturn(CreateListNode(0, 1))).AddTestCase(
	test.NewMethodTestCase(CreateListNode(1, 8), CreateListNode(0)).ExpectReturn(CreateListNode(1, 8)))
