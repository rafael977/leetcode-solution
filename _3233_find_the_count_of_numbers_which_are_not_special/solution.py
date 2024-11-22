# Created by Rafael Shen at 2024/11/22 10:07
# leetgo: 1.4.9
# https://leetcode.cn/problems/find-the-count-of-numbers-which-are-not-special/

from math import sqrt, floor

# @lc code=begin


class Solution:
    def nonSpecialCount(self, l: int, r: int) -> int:
        n = floor(sqrt(r))
        isPrime = [True] * (n + 1)
        ans = r - l + 1
        for i in range(2, n + 1):
            if isPrime[i]:
                if l <= i * i <= r:
                    ans -= 1
                for j in range(i * i, n + 1, i):
                    isPrime[j] = False
        return ans


# @lc code=end
