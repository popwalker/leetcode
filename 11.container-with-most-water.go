/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	var i, j int = 0, len(height)-1
	var res int
	for i != j {
		h := 0
		if height[i] > height[j] {
			h = height[j]
			j--
		} else {
			h = height[i]
			i++
		}
		area := (j-i+1) * h
		if area > res {
			res = area
		}
	}
	return res
}
// @lc code=end

