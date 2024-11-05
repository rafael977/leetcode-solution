# Created by Rafael Shen at 2024/11/05 22:46
# leetgo: 1.4.9
# https://leetcode.cn/problems/find-the-winning-player-in-coin-game/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def losingPlayer(self, x: int, y: int) -> str:
        if x * 4 <= y:
            return "Alice" if x % 2 == 1 else "Bob"
        else:
            return "Alice" if (y // 4) % 2 == 1 else "Bob"


# @lc code=end
