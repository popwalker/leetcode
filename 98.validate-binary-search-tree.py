#
# @lc app=leetcode id=98 lang=python3
#
# [98] Validate Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    # 中序遍历，只保留前驱节点，更节省内存
    def isValidBST(self, root: TreeNode) -> bool:
        self.prev = None
        return self.helper(root)

    def helper(self, root):
        if root is None:
            return True

        if not self.helper(root.left):
            return False

        if self.prev and self.prev.val >= root.val:
            return False

        self.prev = root
        return self.helper(root.right)
# @lc code=end

