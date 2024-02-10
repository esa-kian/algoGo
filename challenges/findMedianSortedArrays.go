/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

 

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/

const (
    MaxInt = int(^uint(0) >> 1)
    MinInt = -MaxInt - 1
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)

    if m > n {
        nums1, nums2 = nums2, nums1
        m, n = n, m
    }
    total := m + n
    half := (total + 1) / 2

    l, r := 0, m-1

    for {
        h1 := (l + r) >> 1
        h2 := half - h1 - 2

        l1 := MinInt
        l2 := MinInt
        r1 := MaxInt
        r2 := MaxInt

        if h1 >= 0 {
            l1 = nums1[h1]
        }
        if h2 >= 0 {
            l2 = nums2[h2]
        }
        if h1+1 < m {
            r1 = nums1[h1+1]
        }
        if h2+1 < n {
            r2 = nums2[h2+1]
        }

        if l1 <= r2 && l2 <= r1 {
            if total%2 == 0 {
                return float64(max(l1, l2)+min(r1, r2)) / 2
            }
            return float64(max(l1, l2))
        }

        if l1 > r2 {
            r = h1 - 1
        } else {
            l = h1 + 1
        }
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
