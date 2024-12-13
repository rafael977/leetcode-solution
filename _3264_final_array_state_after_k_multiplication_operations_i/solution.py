# Created by Rafael Shen at 2024/12/13 10:40
# leetgo: 1.4.11
# https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-i/

import heapq
from typing import List

# @lc code=begin


class Solution:
    def getFinalState(self, nums: List[int], k: int, multiplier: int) -> List[int]:
        h = []
        for i, v in enumerate(nums):
            heapq.heappush(h, (v, i))
        while k > 0:
            _, i = heapq.heappop(h)
            nums[i] *= multiplier
            heapq.heappush(h, (nums[i], i))
            k -= 1
        return nums


# @lc code=end
