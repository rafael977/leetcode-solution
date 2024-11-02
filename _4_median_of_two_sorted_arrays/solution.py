# Created by Rafael Shen at 2024/11/02 15:58
# leetgo: 1.4.9
# https://leetcode.cn/problems/median-of-two-sorted-arrays/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        m, n = len(nums1), len(nums2)

        def findKthMin(k: int) -> int:
            i1, i2 = 0, 0
            while True:
                if i1 == m:
                    return nums2[i2 + k - 1]
                if i2 == n:
                    return nums1[i1 + k - 1]
                if k == 1:
                    return min(nums1[i1], nums2[i2])
                ii1 = min(i1 + k // 2, m) - 1
                ii2 = min(i2 + k // 2, n) - 1
                if nums1[ii1] < nums2[ii2]:
                    k -= ii1 - i1 + 1
                    i1 = ii1 + 1
                else:
                    k -= ii2 - i2 + 1
                    i2 = ii2 + 1

        return (findKthMin((m + n + 1) // 2) + findKthMin((m + n + 2) // 2)) / 2


# @lc code=end
