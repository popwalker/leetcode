/*
 * @lc app=leetcode id=98 lang=java
 *
 * [98] Validate Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    /*
    递归实现方式:
    左子树的最大值要小于root
    右子树的最小值要大于root
    */
    public boolean isValidBST(TreeNode root) {
        if (root == null){
            return true;
        }
        return this.isValid(root, null, null);
    }

    public boolean isValid(TreeNode root, Integer min, Integer max){
        if (root == null) {
            return true;
        }

        if (min != null && root.val <= min){
            return false;
        }
        if (max != null && root.val >= max){
            return false;
        }

        return this.isValid(root.left, min, root.val) && this.isValid(root.right, root.val, max);
    }
}
// @lc code=end

