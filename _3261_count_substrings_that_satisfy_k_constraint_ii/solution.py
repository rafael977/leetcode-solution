# Created by Rafael Shen at 2024/11/13 10:41
# leetgo: 1.4.9
# https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-ii/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def countKConstraintSubstrings(
        self, s: str, k: int, queries: List[List[int]]
    ) -> List[int]:
        n = len(s)
        count = [0, 0]
        prefix = [0] * (n + 1)
        right = [n] * n
        i = 0
        for j in range(n):
            count[int(s[j])] += 1
            while count[0] > k and count[1] > k:
                right[i] = j
                count[int(s[i])] -= 1
                i += 1
            prefix[j + 1] = prefix[j] + j - i + 1

        res = []
        for q in queries:
            l, r = q[0], q[1]
            i = min(right[l], r + 1)
            res.append((i - l + 1) * (i - l) // 2 + prefix[r + 1] - prefix[i])
        return res


# @lc code=end
