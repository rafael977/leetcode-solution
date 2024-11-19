# Created by Rafael Shen at 2024/11/19 23:22
# leetgo: 1.4.9
# https://leetcode.cn/problems/count-and-say/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"
        prev = self.countAndSay(n - 1)
        cc = 0
        num = prev[0]
        res = ""
        for ch in prev:
            if num == ch:
                cc += 1
                continue
            res += f"{cc}{num}"
            cc = 1
            num = ch
        res += f"{cc}{num}"
        return res


# @lc code=end
