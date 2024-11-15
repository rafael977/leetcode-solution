# Created by Rafael Shen at 2024/11/15 11:25
# leetgo: 1.4.9
# https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def minFlips(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        r, c = 0, 0
        for i in range(m):
            j, k = 0, n - 1
            while j < k:
                if grid[i][j] != grid[i][k]:
                    r += 1
                j += 1
                k -= 1
        for j in range(n):
            i, k = 0, m - 1
            while i < k:
                if grid[i][j] != grid[k][j]:
                    c += 1
                i += 1
                k -= 1
        return min(r, c)


# @lc code=end
