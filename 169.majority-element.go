/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

import (
	"sort"
)

// @lc code=start
func majorityElement1(nums []int) int {
	/*
	解题思路：
	1. 双层循环暴力破解O(N^2)
	2. 利用map计数,计数大于 n/2,则找到答案, 时间:O(N), 空间:O(N)
	3. 先排序，然后遍历数组校验每个元素重复的次数,O(NlogN)
	4. 分治，数组拆分，每一部分单独找到重复最大的值,复杂度稍高 O(NlogN)
	*/
	var m = make(map[int]int, 0)
	for _, v := range nums {
		elem, ok := m[v]
		if !ok{
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
		if elem >= len(nums) / 2 {
			return v
		}
	}
	return 0
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i+1 == len(nums){
			break
		}

		if nums[i] == nums[i+1]{
			var count int
			for j := i; j < len(nums); j++{
				if nums[j] == nums[i]{
					count += 1
				} else {
					break
				}
			}
			if count > len(nums) / 2{
				return nums[i]
			}
		}
	}
	return 0

}
// @lc code=end

