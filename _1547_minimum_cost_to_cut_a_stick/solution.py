# Created by Rafael Shen at 2024/11/11 10:41
# leetgo: 1.4.9
# https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/

from functools import cache
from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def minCost(self, n: int, cuts: List[int]) -> int:
        cuts.sort()
        cuts = [0] + cuts + [n]

        @cache
        def cut(i: int, j: int) -> int:
            if i + 1 == j:
                return 0
            return (
                min(cut(i, k) + cut(k, j) for k in range(i + 1, j)) + cuts[j] - cuts[i]
            )

        return cut(0, len(cuts) - 1)


# @lc code=end

if __name__ == "__main__":
    n: int = deserialize("int", read_line())
    cuts: List[int] = deserialize("List[int]", read_line())
    ans = Solution().minCost(n, cuts)
    print("\noutput:", serialize(ans, "integer"))
