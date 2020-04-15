/*
 * @lc app=leetcode id=236 lang=java
 *
 * [236] Lowest Common Ancestor of a Binary Tree
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
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        /*
        思路: 
        1. 寻找两个值的路径，如果有重复，则说明有共同祖先
        2. 递归，分别递归左子树和右子树，节点等于p或q则返回
        */
        
        if (root == null || root == p || root == q){
            return root;
        }

        TreeNode left = lowestCommonAncestor(root.left,  p,  q);
        TreeNode right = lowestCommonAncestor(root.right,  p,  q);

        if (left == null){
            return right;
        }

        if (right == null){
            return left;
        }

        return root;
    }
}
// @lc code=end

