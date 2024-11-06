# Created by Rafael Shen at 2024/11/06 10:35
# leetgo: 1.4.9
# https://leetcode.cn/problems/find-the-power-of-k-size-subarrays-i/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def resultsArray(self, nums: List[int], k: int) -> List[int]:
        i, j = 0, k - 1
        ans = []
        while j < len(nums):
            for k in range(i + 1, j + 1):
                if nums[k - 1] + 1 != nums[k]:
                    ans.append(-1)
                    break
            if len(ans) < i + 1:
                ans.append(nums[j])
            i += 1
            j += 1
        return ans


# @lc code=end
