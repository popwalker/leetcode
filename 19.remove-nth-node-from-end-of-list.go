/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 双指针解法, cur指针向前走N步，cur指针走到链尾时，
	// prev的下一个节点即为需要删除的节点
	var prev, cur = head, head
	for i := 0; i < n; i++ {
		cur = cur.Next
	}

	if cur == nil {
		return prev.Next
	}

	for cur.Next != nil {
		prev, cur = prev.Next, cur.Next
	}
	prev.Next = prev.Next.Next

	return head
}

// @lc code=end

