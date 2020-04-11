#
# @lc app=leetcode id=242 lang=python3
#
# [242] Valid Anagram
#
#思路:
# 1. sort
# 对两个字符串进行排序，排序后相同，则说明anagram
# string排序，最快是快排，时间复杂度为O(nlogn)

# 2. map计数
# 对字符串的每个字符，计算出现的次数，保存在map中，
# 然后再对比两个map中每个key的数量是否相同
# @lc code=start
class Solution:
    # 使用排序方式实现
    def isAnagram1(self, s: str, t: str) -> bool:
        return sorted(s) == sorted(t)

    # 使用hashMap实现
    def isAnagram2(self, s: str, t: str) -> bool:
        dict1, dict2 = {}, {}
        for item in s:
            dict1[item] = dict1.get(item, 0) + 1
        for item in t:
            dict2[item] = dict2.get(item, 0) + 1
        return dict1 == dict2
    # 使用hashMap实现(不使用标准库的map)
    def isAnagram(self, s: str, t: str) -> bool:
        dict1, dict2 = [0]*26, [0]*26
        for item in s:
            dict1[ord(item) - ord('a')] += 1
        for item in t:
            dict2[ord(item) - ord('a')] += 1
        return dict1 == dict2
# @lc code=end

