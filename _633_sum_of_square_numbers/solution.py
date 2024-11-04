# Created by Rafael Shen at 2024/11/04 10:24
# leetgo: 1.4.9
# https://leetcode.cn/problems/sum-of-square-numbers/

# @lc code=begin
import math


class Solution:
    def judgeSquareSum(self, c: int) -> bool:
        a = 0
        while a * a <= c:
            b = math.sqrt(c - a * a)
            if b == int(b):
                return True
            a += 1
        return False


# @lc code=end
