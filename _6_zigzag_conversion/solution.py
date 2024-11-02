# Created by Rafael Shen at 2024/11/02 20:17
# leetgo: 1.4.9
# https://leetcode.cn/problems/zigzag-conversion/

# @lc code=begin


class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        rows = [""] * numRows
        r = 0
        dir = 1
        for c in s:
            rows[r] += c
            if (dir == 1 and r == numRows - 1) or (dir == -1 and r == 0):
                dir = -dir
            r += dir
        return "".join(rows)


# @lc code=end
