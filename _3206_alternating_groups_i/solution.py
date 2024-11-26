# Created by Rafael Shen at 2024/11/26 20:51
# leetgo: 1.4.11
# https://leetcode.cn/problems/alternating-groups-i/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def numberOfAlternatingGroups(self, colors: List[int]) -> int:
        n = len(colors)
        ans = 0
        for i in range(n):
            if (
                colors[i % n] == 0
                and colors[(i + 1) % n] == 1
                and colors[(i + 2) % n] == 0
            ) or (
                colors[i % n] == 1
                and colors[(i + 1) % n] == 0
                and colors[(i + 2) % n] == 1
            ):
                ans += 1
        return ans


# @lc code=end
