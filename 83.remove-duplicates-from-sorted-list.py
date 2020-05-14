#
# @lc app=leetcode id=83 lang=python3
#
# [83] Remove Duplicates from Sorted List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates1(self, head: ListNode) -> ListNode:
        if not head or not head.next:
            return head

        curr = head
        while curr and curr.next:
            if curr.val == curr.next.val:
                curr.next = curr.next.next
            else:
                curr = curr.next
        
        return head
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        '''递归实现
        首先判断是否至少有两个结点，若不是的话，直接返回 head。否则对 head->next 调用递归函数，并赋值给 head->next。这里可能比较晕，先看后面一句，返回的时候，head 结点先跟其身后的结点进行比较，如果值相同，那么返回后面的一个结点，当前的 head 结点就被跳过了，而如果不同的话，还是返回 head 结点。可以发现了，进行实质上的删除操作是在最后一句进行了，再来看第二句，对 head 后面的结点调用递归函数，那么就应该 suppose 返回来的链表就已经没有重复项了，此时接到 head 结点后面，在第三句的时候再来检查一下 head 是否又 duplicate 了，实际上递归一直走到了末尾结点，再不断的回溯回来，进行删除重复结点
        '''
        if not head or not head.next:
            return head
        
        head.next = self.deleteDuplicates(head.next)
        
        if head.val == head.next.val:
            return head.next
        else:
            return head
        

# @lc code=end

