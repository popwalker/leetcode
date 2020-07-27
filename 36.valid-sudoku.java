/*
 * @lc app=leetcode id=36 lang=java
 *
 * [36] Valid Sudoku
 */

// @lc code=start
class Solution {
    public boolean isValidSudoku(char[][] board) {
        if (board == null || board.length == 0) return false;
        return isValidBoard(board);
    }

    public boolean isValidBoard(char[][] board) {
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (board[i][j] != '.') {
                    char c = board[i][j];
                    if (!isValid(board, i, j, c)) {
                        return false;
                    }
                }
            }
        }
        return true;
    }

    public boolean isValid(char[][] board, int row, int col, char c) {
        for (int i = 0; i < 9; i++) {
            if (i != col && board[row][i] != '.' && board[row][i] == c) return false;
            if (i != row && board[i][col] != '.' && board[i][col] == c) return false;
            int subBoxRow = 3 * (row/3) + i /3;
            int subBoxCol = 3 * (col/3) + i % 3;
            if (subBoxRow != row && subBoxCol != col && board[subBoxRow][subBoxCol] != '.' && board[subBoxRow][subBoxCol] == c) return false;
        }
        return true;
    }
}
// @lc code=end

