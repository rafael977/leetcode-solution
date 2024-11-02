# Created by Rafael Shen at 2024/11/02 15:10
# leetgo: 1.4.9
# https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal/

# @lc code=begin


class Solution:
    def minChanges(self, n: int, k: int) -> int:
        bs = n ^ k
        return -1 if bs & k != 0 else bs.bit_count()


# @lc code=end
