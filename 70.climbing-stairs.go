/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
    return climbStairs2(n)
}

func climbStairs2(n int)int{
	// 将空间复杂度降到O(1)
	if n == 1 {
		return 1
	}

	var first = 1
	var second = 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

func climbStairs1(n int)int{
	/*
		动态规划
	*/
	if n == 1 {
		return 1
	}

	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}


// @lc code=end

