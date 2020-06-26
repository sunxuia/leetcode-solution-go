package q001_two_sum

import (
	"../../util/test"
	"testing"
)

func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int)
	println(indexes)
	for i, num := range nums {
		if idx, exists := indexes[target-num]; exists {
			return []int{idx, i}
		}
		indexes[num] = i
	}
	return nil
}

var runner = test.NewMethodTestRunner().AddTestCase(
	test.NewMethodTestCase([]int{2, 7, 11, 15}, 9).ExpectReturn([]int{0, 1}))

func TestTwoSum(t *testing.T) {
	runner.Test(t, twoSum)
}
