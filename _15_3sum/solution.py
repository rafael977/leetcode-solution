# Created by Rafael Shen at 2024/11/06 16:49
# leetgo: 1.4.9
# https://leetcode.cn/problems/3sum/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        ans = set()
        nums.sort()
        for i in range(n - 2):
            x = nums[i]
            if i > 0 and x == nums[i - 1]:
                continue
            if x + nums[i + 1] + nums[i + 2] > 0:
                break
            if x + nums[n - 2] + nums[n - 1] < 0:
                continue
            j, k = i + 1, n - 1
            while j < k:
                if x + nums[j] + nums[k] < 0:
                    j += 1
                elif x + nums[j] + nums[k] > 0:
                    k -= 1
                else:
                    ans.add((x, nums[j], nums[k]))
                    j += 1
                    k -= 1
        return [list(tp) for tp in ans]


# @lc code=end
