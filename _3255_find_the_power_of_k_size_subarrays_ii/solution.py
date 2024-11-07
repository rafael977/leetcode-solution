# Created by Rafael Shen at 2024/11/07 09:52
# leetgo: 1.4.9
# https://leetcode.cn/problems/find-the-power-of-k-size-subarrays-ii/

from typing import *

# @lc code=begin


class Solution:
    def resultsArray(self, nums: List[int], k: int) -> List[int]:
        cnt = 0
        n = len(nums)
        ans = [-1] * (n - k + 1)
        for i in range(n):
            cnt = cnt + 1 if i > 0 and nums[i] - nums[i - 1] == 1 else 1
            if cnt >= k:
                ans[i - k + 1] = nums[i]
        return ans


# @lc code=end
