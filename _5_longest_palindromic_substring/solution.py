# Created by Rafael Shen at 2024/11/02 20:04
# leetgo: 1.4.9
# https://leetcode.cn/problems/longest-palindromic-substring/

# @lc code=begin


class Solution:
    def longestPalindrome(self, s: str) -> str:
        maxLen = 0
        ans = ""
        for i in range(len(s)):
            a, b = i, i
            while a >= 0 and b < len(s) and s[a] == s[b]:
                if maxLen < b - a + 1:
                    maxLen = b - a + 1
                    ans = s[a : b + 1]
                a -= 1
                b += 1
            a, b = i, i + 1
            while a >= 0 and b < len(s) and s[a] == s[b]:
                if maxLen < b - a + 1:
                    maxLen = b - a + 1
                    ans = s[a : b + 1]
                a -= 1
                b += 1
        return ans


# @lc code=end
