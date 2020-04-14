#
# @lc app=leetcode id=21 lang=python3
#
# [21] Merge Two Sorted Lists
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        # 双指针法，两个指针分别指向两个链表，每次比较两个元素的值，小的追加到结果中，指针后移

        # 边界判断
        if not l1:
            return l2
        if not l2:
            return l1
        
        result = ListNode()
        cur = result
        while l1 and l2:
            if l1.val < l2.val:
                cur.next = l1
                l1 = l1.next
            else:
                cur.next = l2
                l2 = l2.next
            cur = cur.next
        
        if l1:
            cur.next = l1
        
        if l2:
            cur.next = l2

        return result.next
# @lc code=end

