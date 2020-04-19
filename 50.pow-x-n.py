#
# @lc app=leetcode id=50 lang=python3
#
# [50] Pow(x, n)
#

# @lc code=start
class Solution:
    '''
    1. 调用库函数pow实现, 时间复杂度: O(1)
    2. 使用循环暴力相乘, 时间复杂度: O(N)
    3. 分治思想，每次将n均分两份，计算x的n/2次方，结果相乘，需要考虑奇数偶数的情况,, 时间复杂度: O(logN)
    当n时偶数，f = x^(n/2), result = f*f
    当n时奇数时, f = x^((n-1)/2), result = f*f*x
    '''

    def myPow1(self, x: float, n: int) -> float:
        if not n:
            return  1
        if n < 0:
            return 1 / self.myPow(x, -n)
        if n % 2:
            return  x * self.myPow(x, n-1)
        return self.myPow(x*x, n/2)

    def myPow(self, x: float, n: int) -> float:
        '''
            循环写法
        '''
        if n < 0:
            x = 1 /x
            n = -n
        pow = 1
        while n:
            if n & 1:
                pow *= x
            x *= x
            n >>= 1
        return pow
# @lc code=end

