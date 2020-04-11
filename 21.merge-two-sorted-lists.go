/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l *ListNode

	for l1.Next != nil {
		e1 := l1.Val
		e2 := l2.Val

		if e1 > e2 {
			l.Next = &ListNode{
				Val: e2,
			}
		} else {
			l.Val = e2
			l.Next = &ListNode{
				Val: e1,
			}
		}

		l1 = l1.Next
		l2 = l2.Next
	}
}

// @lc code=end

