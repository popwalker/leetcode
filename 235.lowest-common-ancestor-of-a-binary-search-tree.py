#
# @lc app=leetcode id=235 lang=python3
#
# [235] Lowest Common Ancestor of a Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:

    # 递归实现
    def lowestCommonAncestorWithRecursion(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if p.val < root.val > q.val:
            return self.lowestCommonAncestorWithRecursion(root.left, p, q)
        elif p.val > root.val < q.val:
            return self.lowestCommonAncestorWithRecursion(root.right, p, q)
        return root
    
    # 循环实现
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        while root:
            if p.val < root.val > q.val:
                root = root.left
            elif p.val > root.val < q.val:
                root = root.right
            else:
                return root
# @lc code=end

