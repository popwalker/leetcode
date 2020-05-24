/*
 * @lc app=leetcode id=122 lang=java
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
class Solution {
    public int maxProfit(int[] prices) {
        int profit = 0;
        for (int i = 0; i<prices.length; i++){
            if (i +1 == prices.length){
                break;
            }
            int currPrice = prices[i];
            int nextPrice = prices[i+1];
            if (currPrice < nextPrice) {
                profit += nextPrice - currPrice;
                continue;
            }
        }
        return profit;
    }
}
// @lc code=end

