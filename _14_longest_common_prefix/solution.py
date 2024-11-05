# Created by Rafael Shen at 2024/11/05 23:06
# leetgo: 1.4.9
# https://leetcode.cn/problems/longest-common-prefix/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""
        strs.sort()
        a, b = strs[0], strs[-1]
        i = 0
        while i < len(a) and i < len(b):
            if a[i] != b[i]:
                break
            i += 1
        return a[:i]


# @lc code=end
