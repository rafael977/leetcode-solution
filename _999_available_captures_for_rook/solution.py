# Created by Rafael Shen at 2024/12/06 10:00
# leetgo: 1.4.11
# https://leetcode.cn/problems/available-captures-for-rook/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def numRookCaptures(self, board: List[List[str]]) -> int:
        n = len(board)

        def check(i, j, d):
            if i < 0 or i >= n or j < 0 or j >= n or board[i][j] == "B":
                return 0
            if board[i][j] == "p":
                return 1
            return check(i + d[0], j + d[1], d)

        ans = 0
        for i in range(n):
            for j in range(n):
                if board[i][j] != "R":
                    continue
                ans += (
                    check(i, j, (-1, 0))
                    + check(i, j, (0, 1))
                    + check(i, j, (1, 0))
                    + check(i, j, (0, -1))
                )
        return ans


# @lc code=end
