# Created by Rafael Shen at 2024/11/03 13:35
# leetgo: 1.4.9
# https://leetcode.cn/problems/string-to-integer-atoi/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def myAtoi(self, s: str) -> int:
        s = s.lstrip()
        if len(s) == 0:
            return 0
        sign = 1
        if s[0] == "-":
            sign = -1
        if s[0] == "-" or s[0] == "+":
            s = s[1:]
        ans = 0
        for c in s:
            if not c.isdigit():
                break
            ans = ans * 10 + int(c)
        ans = ans * sign
        if ans < -(2**31):
            return -(2**31)
        elif ans > 2**31 - 1:
            return 2**31 - 1
        else:
            return ans


# @lc code=end
