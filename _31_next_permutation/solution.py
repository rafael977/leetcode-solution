# Created by Rafael Shen at 2024/11/13 12:32
# leetgo: 1.4.9
# https://leetcode.cn/problems/next-permutation/

from typing import List

# @lc code=begin


class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if len(nums) == 1:
            return nums
        i = len(nums) - 2
        while i >= 0:
            if nums[i + 1] > nums[i]:
                break
            i -= 1
        if i == -1:
            nums.reverse()
            return nums
        j = len(nums) - 1
        while j >= i:
            if nums[j] > nums[i]:
                break
            j -= 1
        nums[i], nums[j] = nums[j], nums[i]
        # reverse i+1-->n-1
        k, l = i + 1, len(nums) - 1
        while k < l:
            nums[k], nums[l] = nums[l], nums[k]
            k += 1
            l -= 1

        return nums


# @lc code=end
