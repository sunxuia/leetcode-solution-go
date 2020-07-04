package q023_merge_k_sorted_lists

import (
	"container/heap"
	"testing"

	. "github.com/sunxuia/leetcode-solution-go/src/util/provided"
)

func TestMergeKLists(t *testing.T) {
	doTest(t, mergeKLists)
}

// 时间复杂度 O(NlogK), N  = 所有节点数量, K = len(lists)
func mergeKLists(lists []*ListNode) *ListNode {
	pq := &priorityQueue{}
	heap.Init(pq)
	for _, v := range lists {
		if v != nil {
			heap.Push(pq, v)
		}
	}

	dummy := &ListNode{0, nil}
	curr := dummy
	for len(*pq) > 0 {
		next := heap.Pop(pq).(*ListNode)
		curr.Next = next
		curr = next
		if next.Next != nil {
			heap.Push(pq, next.Next)
		}
	}
	return dummy.Next
}

type priorityQueue []*ListNode

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}
func (pq *priorityQueue) Pop() interface{} {
	val := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return val
}
