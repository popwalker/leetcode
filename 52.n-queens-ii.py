#
# @lc app=leetcode id=52 lang=python3
#
# [52] N-Queens II
#

# @lc code=start
class Solution:
    def totalNQueens(self, n: int) -> int:
        self.result = []
        self.col, self.pie, self.na = set(), set(), set()
        self.DFS(n, 0, [])
        return len(self.result)
    
    def DFS(self, n, row, current_state):
        if row >= n:
            self.result.append(current_state)
            return
        
        for col in range(n):
            if col in self.col or row + col in self.pie or row - col in self.na:
                continue

            self.col.add(col)
            self.pie.add(row + col)
            self.na.add(row - col)

            self.DFS(n, row + 1, current_state + [col])

            self.col.remove(col)
            self.pie.remove(row + col)
            self.na.remove(row - col)
# @lc code=end

