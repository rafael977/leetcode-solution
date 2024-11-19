# Created by Rafael Shen at 2024/11/19 22:49
# leetgo: 1.4.9
# https://leetcode.cn/problems/sudoku-solver/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        row = [0] * 9
        col = [0] * 9
        square = [[0] * 3 for _ in range(3)]
        pos = []
        valid = False
        for i in range(9):
            for j in range(9):
                if board[i][j] == ".":
                    pos.append((i, j))
                else:
                    x = 1 << int(board[i][j])
                    row[i] |= x
                    col[j] |= x
                    square[i // 3][j // 3] |= x

        def dfs(p):
            nonlocal valid
            if p == len(pos):
                valid = True
                return

            i, j = pos[p]
            for x in range(1, 10):
                mask = 1 << x
                if (
                    row[i] & mask == 0
                    and col[j] & mask == 0
                    and square[i // 3][j // 3] & mask == 0
                ):
                    board[i][j] = str(x)
                    row[i] |= mask
                    col[j] |= mask
                    square[i // 3][j // 3] |= mask
                    dfs(p + 1)
                    row[i] ^= mask
                    col[j] ^= mask
                    square[i // 3][j // 3] ^= mask
                if valid:
                    return

        dfs(0)


# @lc code=end
