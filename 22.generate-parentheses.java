import java.util.ArrayList;

/*
 * @lc app=leetcode id=22 lang=java
 *
 * [22] Generate Parentheses
 */

// @lc code=start
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> list = new ArrayList();
        _gen(list, 0, 0, n, "");
        return list;
    }

    public void _gen(List<String> list, int left, int right, int n, String current){
        if (left == n && right == n){
            list.add(current);
            return;
        }

        if (left < n){
            _gen(list, left+1, right, n, current+"(");
        }
        if (right < n && left > right){
            _gen(list, left, right+1, n, current+")");
        }
    }
}
// @lc code=end

