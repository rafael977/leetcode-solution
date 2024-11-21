# Created by Rafael Shen at 2024/11/21 10:10
# leetgo: 1.4.9
# https://leetcode.cn/problems/snake-in-matrix/

from typing import List

# @lc code=begin


class Solution:
    def finalPositionOfSnake(self, n: int, commands: List[str]) -> int:
        pos = 0
        for c in commands:
            match c:
                case "UP":
                    pos -= n
                case "DOWN":
                    pos += n
                case "LEFT":
                    pos -= 1
                case "RIGHT":
                    pos += 1
        return pos


# @lc code=end
