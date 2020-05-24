/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	/*
	解题方法:
	1. DFS 深度优先搜索算法
	2. 贪心算法
	3. DF 动态规划

	本题使用贪心算法进行解题
	*/
	var profit int
	for i := 0; i < len(prices); i++ {
		if i+1 == len(prices){
			break
		}
		currP := prices[i]
		nextP := prices[i+1]
		if currP < nextP {
			profit += nextP-currP
			continue
		}
	}
	return profit
}
// @lc code=end

