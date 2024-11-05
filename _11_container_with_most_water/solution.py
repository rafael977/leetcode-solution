# Created by Rafael Shen at 2024/11/04 15:09
# leetgo: 1.4.9
# https://leetcode.cn/problems/container-with-most-water/

from typing import *

# @lc code=begin


class Solution:
    def maxArea(self, height: List[int]) -> int:
        i, j = 0, len(height) - 1
        ans = 0
        while i < j:
            ans = max(ans, min(height[i], height[j]) * (j - i))
            if height[i] < height[j]:
                i += 1
            else:
                j -= 1
        return ans


# @lc code=end
