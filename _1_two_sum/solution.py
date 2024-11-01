# Created by Rafael Shen at 2024/11/01 15:29
# leetgo: 1.4.9
# https://leetcode.cn/problems/two-sum/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        im = dict()
        for i in range(0, len(nums)):
            if nums[i] in im:
                return [im[nums[i]], i]
            im[target - nums[i]] = i
        return []


# @lc code=end
