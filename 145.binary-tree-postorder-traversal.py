#
# @lc app=leetcode id=145 lang=python3
#
# [145] Binary Tree Postorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        return self.postorderIterate(root)

    def postorderIterate(self, root):
        """
        迭代实现
        """

        if not root: return []

        stack, res = [root], []
        while stack:
            node = stack.pop()
            if node:
                stack.append(node)
                stack.append(None)
                if node.right: stack.append(node.right)
                if node.left: stack.append(node.left)
            else:
                v = stack.pop()
                res.append(v.val)
        return res




    def postorderTraversal1(self, root):
        """
        递归实现
        """
        res = []
        self.postorderRecrusive(root, res)
        return res

    def postorderRecrusive(self, root, res):
        if root:
            self.postorderRecrusive(root.left, res)
            self.postorderRecrusive(root.right, res)
            res.append(root.val)
# @lc code=end

