/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

 // 递归解法
 func lowestCommonAncestorWithRecursion(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val{
		return lowestCommonAncestorWithRecursion(root.Left, p, q)
	} else if (p.Val > root.Val && q.Val > root.Val){
		return lowestCommonAncestorWithRecursion(root.Right, p, q)
	}

	return root
}

// 循环解法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val{
			root = root.Left
		} else if (p.Val > root.Val && q.Val > root.Val){
			root = root.Right
		} else {
			return root
		}
	}
	return root
}
// @lc code=end

