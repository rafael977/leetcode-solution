# Created by Rafael Shen at 2024/11/03 13:55
# leetgo: 1.4.9
# https://leetcode.cn/problems/palindrome-number/

# @lc code=begin


class Solution:
    def isPalindrome(self, x: int) -> bool:
        sx = f"{x}"
        return sx == sx[::-1]


# @lc code=end
