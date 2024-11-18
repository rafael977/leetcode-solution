# Created by Rafael Shen at 2024/11/18 21:04
# leetgo: 1.4.9
# https://leetcode.cn/problems/search-insert-position/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        i, j = 0, len(nums)
        while i < j:
            mid = (i + j) // 2
            if nums[mid] < target:
                i = mid + 1
            else:
                j = mid
        return i


# @lc code=end
