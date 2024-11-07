# Created by Rafael Shen at 2024/11/07 16:07
# leetgo: 1.4.9
# https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def __init__(self):
        self.map = {
            "2": ["a", "b", "c"],
            "3": ["d", "e", "f"],
            "4": ["g", "h", "i"],
            "5": ["j", "k", "l"],
            "6": ["m", "n", "o"],
            "7": ["p", "q", "r", "s"],
            "8": ["t", "u", "v"],
            "9": ["w", "x", "y", "z"],
        }

    def letterCombinations(self, digits: str) -> List[str]:
        if digits == "":
            return []
        ans = [""]
        for c in digits:
            ans = [a + b for a in ans for b in self.map[c]]
        return ans


# @lc code=end

if __name__ == "__main__":
    digits: str = deserialize("str", read_line())
    ans = Solution().letterCombinations(digits)
    print("\noutput:", serialize(ans, "string[]"))
