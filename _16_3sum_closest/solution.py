# Created by Rafael Shen at 2024/11/06 17:08
# leetgo: 1.4.9
# https://leetcode.cn/problems/3sum-closest/

from typing import *

# @lc code=begin


class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        n = len(nums)
        nums.sort()
        ans = 0
        d = 1e5
        for i in range(n - 2):
            j, k = i + 1, n - 1
            while j < k:
                sum = nums[i] + nums[j] + nums[k]
                if abs(sum - target) < d:
                    ans = sum
                    d = abs(sum - target)
                if sum <= target:
                    j += 1
                elif sum > target:
                    k -= 1
        return ans


# @lc code=end
