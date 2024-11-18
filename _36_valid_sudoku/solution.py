# Created by Rafael Shen at 2024/11/18 21:11
# leetgo: 1.4.9
# https://leetcode.cn/problems/valid-sudoku/

from typing import List

# @lc code=begin


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row = [0 for _ in range(9)]
        col = [0 for _ in range(9)]
        square = [0 for _ in range(9)]

        for i in range(9):
            for j in range(9):
                if board[i][j] == ".":
                    continue
                x = int(board[i][j])
                mask = 1 << x
                if (
                    row[i] & mask != 0
                    or col[j] & mask != 0
                    or square[i // 3 * 3 + j // 3] & mask != 0
                ):
                    return False
                row[i] = row[i] | mask
                col[j] = col[j] | mask
                square[i // 3 * 3 + j // 3] = square[i // 3 * 3 + j // 3] | mask

        return True


# @lc code=end
