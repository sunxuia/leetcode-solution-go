package q041_first_missing_positive

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	doTest(t, firstMissingPositive)
}

// 因为最小的缺失整数肯定 <= len(nums),
// 所以将数字val 放到 nums[val-1] 的位置上,
// 这样最小的整数值肯定是 nums[i] != i + 1
func firstMissingPositive(nums []int) int {
	for _, val := range nums {
		for 0 < val && val <= len(nums) && nums[val-1] != val {
			val, nums[val-1] = nums[val-1], val
		}
	}
	for i, val := range nums {
		if i+1 != val {
			return i + 1
		}
	}
	return len(nums) + 1
}
