#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        dict = {}
        for i,x in enumerate(nums):
            y = target - x
            j = dict.get(y)
            if j is not None:
                return [j, i]
            dict[x]=i
# @lc code=end

