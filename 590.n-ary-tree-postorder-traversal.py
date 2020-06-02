#
# @lc app=leetcode id=590 lang=python3
#
# [590] N-ary Tree Postorder Traversal
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def postorder(self, root: 'Node') -> List[int]:
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
                stack.extend(node.children[::-1])
            else:
                v = stack.pop()
                res.append(v.val)
        return res
# @lc code=end

