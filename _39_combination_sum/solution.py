# Created by Rafael Shen at 2024/11/21 13:52
# leetgo: 1.4.9
# https://leetcode.cn/problems/combination-sum/

from typing import List

# @lc code=begin


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        picked = []
        ans = []

        def backtrack(i: int, t: int):
            if t == 0:
                ans.append(list(picked))
                return
            for j in range(i, len(candidates)):
                c = candidates[j]
                if t - c < 0:
                    break
                picked.append(c)
                backtrack(j, t - c)
                picked.pop()

        backtrack(0, target)
        return ans


# @lc code=end
