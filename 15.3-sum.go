package leetcode

import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	/*
		solution: sort & find
	*/
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]

			if s < 0 {
				l += 1
			} else if s > 0 {
				r -= 1
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l += 1
				r -= 1
			}
		}
	}
	return result
}

// @lc code=end
