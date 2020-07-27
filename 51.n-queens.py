#
# @lc app=leetcode id=51 lang=python3
#
# [51] N-Queens
#

# @lc code=start
class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        self.result = []
        self.col, self.pie, self.na = set(), set(), set()
        self.DFS(n, 0, [])
        return self._generate_result(n)
    
    def DFS(self, n, row, current_result):
        if row >= n:
            self.result.append(current_result)
            return
        
        for col in range(n):
            if col in self.col or col + row in self.pie or row - col in self.na:
                continue

            self.col.add(col)
            self.pie.add(col+row)
            self.na.add(row-col)

            self.DFS(n, row+1, current_result+[col])

            self.col.remove(col)
            self.pie.remove(col+row)
            self.na.remove(row-col)
    def _generate_result(self, n):
        board = []
        for res in self.result:
            for i in res:
                board.append("."*i + "Q" + "."*(n-i-1))
        return [board[i: i+n] for i in range(0, len(board), n)]
# @lc code=end

