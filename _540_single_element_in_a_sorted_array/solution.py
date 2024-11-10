# Created by Rafael Shen at 2024/11/10 13:20
# leetgo: 1.4.9
# https://leetcode.cn/problems/single-element-in-a-sorted-array/

from typing import List

# @lc code=begin


class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        lo, hi = 0, len(nums) - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] == nums[mid ^ 1]:
                lo = mid + 1
            else:
                hi = mid
        return nums[lo]


# @lc code=end
