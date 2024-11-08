# Created by Rafael Shen at 2024/11/08 12:16
# leetgo: 1.4.9
# https://leetcode.cn/problems/valid-parentheses/

from typing import *

# @lc code=begin


class Solution:
    def __init__(self):
        self.pairs = {
            ")": "(",
            "]": "[",
            "}": "{",
        }

    def isValid(self, s: str) -> bool:
        stack = list()
        for c in s:
            if c in self.pairs:
                if not stack or stack[-1] != self.pairs[c]:
                    return False
                stack.pop()
                continue
            stack.append(c)
        return len(stack) == 0


# @lc code=end
