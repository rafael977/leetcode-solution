# Created by Rafael Shen at 2024/11/04 11:57
# leetgo: 1.4.9
# https://leetcode.cn/problems/reverse-integer/

# @lc code=begin


class Solution:
    def reverse(self, x: int) -> int:
        INT_MIN, INT_MAX = -(2**31), 2**31 - 1
        sign = 1
        if x < 0:
            sign = -1
            x = -x
        ans = 0
        while x > 0:
            if ans * sign < INT_MIN // 10 + 1 or ans * sign > INT_MAX // 10:
                return 0
            ans = ans * 10 + x % 10
            x = x // 10
        return ans * sign


# @lc code=end
