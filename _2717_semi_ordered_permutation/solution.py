# Created by Rafael Shen at 2024/12/11 10:44
# leetgo: 1.4.11
# https://leetcode.cn/problems/semi-ordered-permutation/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def semiOrderedPermutation(self, nums: List[int]) -> int:
        n = len(nums)
        i, j = 0, n - 1
        for x in range(n):
            if nums[x] == 1:
                i = x
            if nums[x] == n:
                j = x
        ans = i + n - 1 - j
        return ans if i < j else ans - 1


# @lc code=end
