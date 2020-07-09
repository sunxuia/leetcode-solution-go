package q031_next_permutation

import (
	"testing"
)

func TestNextPermutation(t *testing.T) {
	doTest(t, nextPermutation)
}

func nextPermutation(nums []int) {
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		// 从后往前找, 如果左边 < 右边, 说明可以在[i, n-1] 这个范围内进行一次翻转
		if nums[i] < nums[i+1] {
			// 从右边找出1 个比左边值小的放到前面,
			// 剩下的进行反序(因为是倒序的), 得出最小值
			j := n - 1
			for i < j && nums[i] >= nums[j] {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
			reverse(nums, i+1, n-1)
			return
		}
	}
	reverse(nums, 0, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
