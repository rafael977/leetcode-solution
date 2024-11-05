# Created by Rafael Shen at 2024/11/04 15:18
# leetgo: 1.4.9
# https://leetcode.cn/problems/integer-to-roman/

# @lc code=begin


class Solution:
    def __init__(self):
        self.map = {
            1000: "M",
            900: "CM",
            500: "D",
            400: "CD",
            100: "C",
            90: "XC",
            50: "L",
            40: "XL",
            10: "X",
            9: "IX",
            5: "V",
            4: "IV",
            1: "I",
        }

    def intToRoman(self, num: int) -> str:
        ans = ""
        for k, v in self.map.items():
            while num >= k:
                ans += v
                num -= k
        return ans


# @lc code=end
