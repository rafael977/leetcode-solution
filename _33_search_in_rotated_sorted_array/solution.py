# Created by Rafael Shen at 2024/11/15 13:39
# leetgo: 1.4.9
# https://leetcode.cn/problems/search-in-rotated-sorted-array/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        n = len(nums)

        def find_k():
            i, j = 0, n - 1
            while i < j:
                mid = (i + j) // 2
                if mid > 0 and nums[mid] < nums[mid - 1]:
                    return mid
                if nums[mid] > nums[mid + 1]:
                    return mid + 1
                if nums[mid] < nums[n - 1]:
                    j = mid
                elif nums[mid] > nums[0]:
                    i = mid
            return i

        k = find_k()
        i, j = k, n + k - 1
        while i <= j:
            mid = (i + j) // 2
            idx = mid % n
            if nums[idx] == target:
                return idx
            if nums[idx] < target:
                i = mid + 1
            else:
                j = mid - 1
        return -1


# @lc code=end

if __name__ == "__main__":
    nums: List[int] = deserialize("List[int]", read_line())
    target: int = deserialize("int", read_line())
    ans = Solution().search(nums, target)
    print("\noutput:", serialize(ans, "integer"))
