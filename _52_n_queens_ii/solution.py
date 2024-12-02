# Created by Rafael Shen at 2024/12/02 11:30
# leetgo: 1.4.11
# https://leetcode.cn/problems/n-queens-ii/

# @lc code=begin


class Solution:
    def totalNQueens(self, n: int) -> int:
        pos = []
        ans = 0

        def backtrack(r: int):
            nonlocal ans
            if r == n:
                ans += 1
                return
            for c in range(n):
                check = False
                for p in pos:
                    if p[0] == r or p[1] == c or abs(r - p[0]) == abs(c - p[1]):
                        check = True
                        break
                if not check:
                    pos.append((r, c))
                    backtrack(r + 1)
                    pos.pop()

        backtrack(0)
        return ans


# @lc code=end
