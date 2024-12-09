# Created by Rafael Shen at 2024/12/09 10:37
# leetgo: 1.4.11
# https://leetcode.cn/problems/determine-color-of-a-chessboard-square/

# @lc code=begin


class Solution:
    def squareIsWhite(self, coordinates: str) -> bool:
        return ((ord(coordinates[0]) - ord("a")) % 2) ^ (
            (int(coordinates[1]) - 1) % 2
        ) == 1


# @lc code=end
