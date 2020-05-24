#
# @lc app=leetcode id=122 lang=python3
#
# [122] Best Time to Buy and Sell Stock II
#

# @lc code=start
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0
        for i, v in enumerate(prices):
            if i + 1 == len(prices):
                break
            curr_price = v
            next_price =  prices[i+1]
            if curr_price < next_price:
                profit += next_price - curr_price
                continue

        return profit   

        
# @lc code=end

