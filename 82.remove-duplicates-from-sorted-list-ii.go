/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	
	if head.Next != nil &&  head.Val == head.Next.Val{
		for head.Next != nil && head.Val == head.Next.Val{
			head = head.Next
		}
		return deleteDuplicates(head.Next)
		
	}

	head.Next = deleteDuplicates(head.Next)
	return head
}


// @lc code=end

