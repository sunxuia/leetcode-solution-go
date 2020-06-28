package q004_median_of_two_sorted_arrays

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	doTest(t, findMedianSortedArrays)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		nums1, nums2 = nums2, nums1
		len1, len2 = len2, len1
	}

	start, end := 0, len1
	for start <= end {
		i := (start + end) / 2
		j := (len1+len2+1)/2 - i
		if i < end && nums2[j-1] > nums1[i] {
			// 左边界最大值 > 右边界最小值, 且i 还有前进空间.
			start = i + 1
		} else if i > start && nums1[i-1] > nums2[j] {
			// 左边界最大值 > 右边界最小值, 且i 还有后退空间.
			end = i - 1
		} else {
			// 边界确定
			// 找到左边界的最大值(如果总数为奇数的话, 就是中间的中位数)
			var maxLeft float64
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else if nums1[i-1] >= nums2[j-1] {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = float64(nums2[j-1])
			}
			if (len1+len2)%2 == 1 {
				return maxLeft
			}

			// 找到右边界的最小值.
			var minRight float64
			if i == len1 {
				minRight = float64(nums2[j])
			} else if j == len2 {
				minRight = float64(nums1[i])
			} else if nums1[i] <= nums2[j] {
				minRight = float64(nums1[i])
			} else {
				minRight = float64(nums2[j])
			}

			return (maxLeft + minRight) / 2.0
		}
	}
	return 0
}
