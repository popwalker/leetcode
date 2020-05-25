/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var list []string
	_gen(&list, 0, 0, n, "")
	return list
}

func _gen(list *[]string, left, right, n int, current string){
	if left == n && right == n {
		list = append(*list, current)
		return 
	}

	if left < n {
		_gen(list, left + 1, right, n, current+"(")
	}
	if right < n && left > right {
		_gen(list, left, right+1, n, current+")")
	}
}
// @lc code=end

