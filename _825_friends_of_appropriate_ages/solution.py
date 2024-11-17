# Created by Rafael Shen at 2024/11/17 21:16
# leetgo: 1.4.9
# https://leetcode.cn/problems/friends-of-appropriate-ages/

from typing import List

# @lc code=begin


class Solution:
    def numFriendRequests(self, ages: List[int]) -> int:
        ages.sort()
        l, r, ans = 0, 0, 0
        for age in ages:
            if age < 15:
                continue
            while ages[l] <= 0.5 * age + 7:
                l += 1
            while r + 1 < len(ages) and ages[r + 1] <= age:
                r += 1
            ans += r - l
        return ans


# @lc code=end
