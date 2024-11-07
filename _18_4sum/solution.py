# Created by Rafael Shen at 2024/11/07 16:18
# leetgo: 1.4.9
# https://leetcode.cn/problems/4sum/

from typing import *

# @lc code=begin


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        n = len(nums)
        nums.sort()
        ans = set()
        for i in range(n - 3):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            if nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target:
                break
            for j in range(i + 1, n - 2):
                k, l = j + 1, n - 1
                while k < l:
                    if nums[i] + nums[j] + nums[k] + nums[l] < target:
                        k += 1
                    elif nums[i] + nums[j] + nums[k] + nums[l] > target:
                        l -= 1
                    else:
                        ans.add((nums[i], nums[j], nums[k], nums[l]))
                        k += 1
                        l -= 1
        return [list(a) for a in ans]


# @lc code=end
