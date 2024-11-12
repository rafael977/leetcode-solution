# Created by Rafael Shen at 2024/11/12 21:59
# leetgo: 1.4.9
# https://leetcode.cn/problems/divide-two-integers/

# @lc code=begin


class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        def div(a: int, b: int) -> int:
            if a < b:
                return 0
            tb = b
            cnt = 1
            while tb + tb <= a:
                tb = tb + tb
                cnt = cnt + cnt
            return cnt + div(a - tb, b)

        sign = (
            1
            if (dividend > 0 and divisor > 0) or (dividend < 0 and divisor < 0)
            else -1
        )
        dividend = abs(dividend)
        divisor = abs(divisor)
        q = sign * div(dividend, divisor)
        if q < -(2**31):
            return -(2**31)
        if q > 2**31 - 1:
            return 2**31 - 1
        return q


# @lc code=end
