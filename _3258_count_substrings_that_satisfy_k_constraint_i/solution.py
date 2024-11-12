# Created by Rafael Shen at 2024/11/12 21:38
# leetgo: 1.4.9
# https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/

from functools import cache

# @lc code=begin


class Solution:
    def countKConstraintSubstrings(self, s: str, k: int) -> int:
        @cache
        def calcLen(i: int) -> int:
            return (i + 1) * i // 2

        ans = 0
        l, r, pr = 0, 0, 0
        c0, c1 = 0, 0
        while r < len(s):
            c0 += 1 if s[r] == "0" else 0
            c1 += 1 if s[r] == "1" else 0
            if c0 > k or c1 > k:
                ans += calcLen(r - l) - calcLen(pr - l)
                pr = r
                while c0 > k and c1 > k:
                    c0 -= 1 if s[l] == "0" else 0
                    c1 -= 1 if s[l] == "1" else 0
                    l += 1
            r += 1

        ans += calcLen(r - l) - calcLen(pr - l)
        return ans


# @lc code=end
