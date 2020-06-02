#
# @lc app=leetcode id=144 lang=python3
#
# [144] Binary Tree Preorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    """
    递归实现
    """
    def preorderTraversal1(self, root: TreeNode) -> List[int]:
        result = []
        self.preorder(root, result)
        return result

    def preorder(self, root: TreeNode, result: List[int]):
        if root:
            result.append(root.val)
            self.preorder(root.left, result)
            self.preorder(root.right, result)
    
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        """
            迭代实现
        """

        if not root: return []

        stack, res = [root,], []

        while stack:
            node = stack.pop()
            if node:
                if node.right: stack.append(node.right)
                if node.left: stack.append(node.left)
                stack.append(node)
                stack.append(None)
            else:
                v = stack.pop()
                res.append(v.val)

        return res

# @lc code=end

