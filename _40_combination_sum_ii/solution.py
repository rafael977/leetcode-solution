# Created by Rafael Shen at 2024/11/21 14:12
# leetgo: 1.4.9
# https://leetcode.cn/problems/combination-sum-ii/

from typing import List

# @lc code=begin


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        picked = []
        ans = []

        def backtrack(i: int, t: int):
            if t == 0:
                ans.append(list(picked))
                return
            for j in range(i, len(candidates)):
                if j > i and candidates[j] == candidates[j - 1]:
                    continue
                c = candidates[j]
                if t - c < 0:
                    break
                picked.append(c)
                backtrack(j + 1, t - c)
                picked.pop()

        backtrack(0, target)
        return ans


# @lc code=end
