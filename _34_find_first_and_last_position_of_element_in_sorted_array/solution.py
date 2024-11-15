# Created by Rafael Shen at 2024/11/15 14:21
# leetgo: 1.4.9
# https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

from typing import List

# @lc code=begin


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        def search(x):
            i, j = 0, len(nums)
            while i < j:
                mid = (i + j) // 2
                if nums[mid] < x:
                    i = mid + 1
                else:
                    j = mid
            return i

        i = search(target)
        if i == len(nums) or nums[i] != target:
            return [-1, -1]
        return [i, search(target + 1) - 1]


# @lc code=end
