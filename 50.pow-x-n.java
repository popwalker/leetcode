/*
 * @lc app=leetcode id=50 lang=java
 *
 * [50] Pow(x, n)
 */

// @lc code=start
class Solution {
    public double myPow(double x, int n) {
        if (n == 0) {
            return 1;
        }

        if (n < 0) {
            return 1 / myPow(x , -n);
        }

        if (n % 2 == 1){
            return x * myPow(x, n-1);
        }

        return myPow(x*x, n/2);
    }
}
// @lc code=end

