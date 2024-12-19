# Created by Rafael Shen at 2024/12/19 18:10
# leetgo: 1.4.11
# https://leetcode.cn/problems/find-indices-of-stable-mountains/

from typing import List

# @lc code=begin


class Solution:
    def stableMountains(self, height: List[int], threshold: int) -> List[int]:
        ans = []
        for i in range(1, len(height)):
            if height[i - 1] > threshold:
                ans.append(i)
        return ans


# @lc code=end
