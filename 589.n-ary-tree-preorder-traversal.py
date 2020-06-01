#
# @lc app=leetcode id=589 lang=python3
#
# [589] N-ary Tree Preorder Traversal
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
    """
    递归实现
    """
    def preorder1(self, root: 'Node') -> List[int]:
        result = []
        self.help(root, result)
        return result

    def help(self, root: 'Node', result: List[int]):
        if root:
            result.append(root.val)
            for node in root.children:
                self.help(node, result)

    """
    迭代实现
    """
    def preorder(self, root: 'Node') -> List[int]:
        if not root: return []

        stack, res = [root,], []

        while stack:
            node = stack.pop()
            res.append(node.val)
            stack.extend(node.children[::-1])
        return res


# @lc code=end

