package leetcode

import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 递归解法
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min {
		return false
	}

	if root.Val >= max {
		return false
	}

	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}

// @lc code=end
