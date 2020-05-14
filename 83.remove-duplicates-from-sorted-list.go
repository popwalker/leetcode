/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var curr = head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 递归实现
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = deleteDuplicates(head.Next)

	if head.Val == head.Next.Val{
		return head.Next
	}else {
		return head
	}
}

// @lc code=end

