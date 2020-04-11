package leetcode

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram1(s string, t string) bool {
	m1, m2 := make(map[rune]int), make(map[rune]int)
	for _, c := range s {
		m1[c] += 1
	}
	for _, c := range t {
		m2[c] += 1
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}
func isAnagram(s string, t string) bool {
	m1, m2 := [26]int{}, [26]int{}
	for _, c := range s {
		m1[c-'a'] += 1
	}
	for _, c := range t {
		m2[c-'a'] += 1
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2 := m2[k]
		if v1 != v2 {
			return false
		}
	}
	return true
}

// @lc code=end
