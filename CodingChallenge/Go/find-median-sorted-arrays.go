package main

/**
 * Find the median of two sorted arrays.
 * Input: nums1 =[1,3], nums2 =[2]
 * Expected: 2.00000
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge two sorted arrays
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	merged = append(merged, nums1[i:]...)
	merged = append(merged, nums2[j:]...)
	// find median
	totalLen := len(merged)
	if totalLen%2 != 0 {
		return float64(merged[totalLen/2])
	}
	mid1 := merged[(totalLen/2)-1]
	mid2 := merged[totalLen/2]
	return float64(mid1+mid2) / 2.0
}
