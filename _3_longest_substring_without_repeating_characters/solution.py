# Created by Rafael Shen at 2024/11/01 17:08
# leetgo: 1.4.9
# https://leetcode.cn/problems/longest-substring-without-repeating-characters/

from typing import *

# @lc code=begin


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        ans = 0
        b = 0
        m = dict()
        for e in range(len(s)):
            if s[e] not in m or m[s[e]] is False:
                ans = max(ans, e - b + 1)
                m[s[e]] = True
                continue
            while b < e and s[b] != s[e]:
                m[s[b]] = False
                b += 1
            b += 1
            ans = max(ans, e - b + 1)
        return ans


# @lc code=end
