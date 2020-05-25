#
# @lc app=leetcode id=22 lang=python3
#
# [22] Generate Parentheses
#

# @lc code=start
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        self.list = []
        self._gen(0, 0, n, "")
        return self.list
    
    def _gen(self, left, right, n, current):
        """
        left: 左括号已经使用的数量
        right: 有括号已经使用的数量
        n: 总数量
        current: 当前结果
        """
        if left == n and right == n:
            self.list.append(current)
            return
        
        if left < n:
            self._gen(left+1, right, n, current + "(")

        if right < n and left > right:
            self._gen(left, right+1, n, current + ")")

# @lc code=end

