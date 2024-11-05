# Created by Rafael Shen at 2024/11/05 22:58
# leetgo: 1.4.9
# https://leetcode.cn/problems/roman-to-integer/

# @lc code=begin


class Solution:
    def __init__(self):
        self.map = {
            "I": 1,
            "IV": 4,
            "V": 5,
            "IX": 9,
            "X": 10,
            "XL": 40,
            "L": 50,
            "XC": 90,
            "C": 100,
            "CD": 400,
            "D": 500,
            "CM": 900,
            "M": 1000,
        }

    def romanToInt(self, s: str) -> int:
        i = 0
        ans = 0
        while i < len(s):
            if s[i : i + 2] in self.map:
                ans += self.map[s[i : i + 2]]
                i += 2
            else:
                ans += self.map[s[i : i + 1]]
                i += 1
        return ans


# @lc code=end
