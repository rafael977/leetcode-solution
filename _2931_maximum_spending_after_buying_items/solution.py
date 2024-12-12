# Created by Rafael Shen at 2024/12/12 10:24
# leetgo: 1.4.11
# https://leetcode.cn/problems/maximum-spending-after-buying-items/

import heapq
from typing import List

# @lc code=begin


class Solution:
    def maxSpending(self, values: List[List[int]]) -> int:
        m, n = len(values), len(values[0])
        h = []
        for i in range(m):
            heapq.heappush(h, (-values[i][0], i, 0))
        d = m * n
        ans = 0
        while len(h) > 0:
            _, i, j = heapq.heappop(h)
            ans += values[i][j] * d
            d -= 1
            if j + 1 < n:
                heapq.heappush(h, (-values[i][j + 1], i, j + 1))
        return ans


# @lc code=end
