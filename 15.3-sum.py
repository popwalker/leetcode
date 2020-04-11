#
# @lc app=leetcode id=15 lang=python3
#
# [15] 3Sum
#

# @lc code=start
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        if len(nums) < 3:
            return []

        nums.sort()
        res = set()

        for i, a in enumerate(nums[:-2]):
            if i >= 1 and a == nums[i-1]:
                continue

            d = {}
            for b in nums[i+1:]:
                c = -a-b
                if b not in d:
                    d[c] = 1
                else:
                    res.add((a,c,b))

        return map(list, res)

# @lc code=end

