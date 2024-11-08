# Created by Rafael Shen at 2024/11/08 12:47
# leetgo: 1.4.9
# https://leetcode.cn/problems/generate-parentheses/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        def backtrack(nn, cnt):
            if nn == 0:
                ans.append("".join(s) + ")" * cnt)
                return
            s.append("(")
            backtrack(nn - 1, cnt + 1)
            s.pop()
            if cnt > 0:
                s.append(")")
                backtrack(nn, cnt - 1)
                s.pop()

        ans = []
        s = []
        backtrack(n, 0)
        return ans


# @lc code=end

if __name__ == "__main__":
    n: int = deserialize("int", read_line())
    ans = Solution().generateParenthesis(n)
    print("\noutput:", serialize(ans, "string[]"))
