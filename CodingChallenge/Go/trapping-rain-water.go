package main

/**
 * Trapping Rain Water
 * Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 */
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	water := 0
	for left < right {
		if height[left] < height[right] {
			left++
			leftMax = max(leftMax, height[left])
			water += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			water += rightMax - height[right]
		}
	}
	return water
}
