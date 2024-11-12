# Created by Rafael Shen at 2024/11/12 22:49
# leetgo: 1.4.9
# https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        wl = len(words[0])
        wordDict = {}
        for w in words:
            wordDict[w] = wordDict[w] + 1 if w in wordDict else 1

        def allZero(wc: dict) -> bool:
            cnt = 0
            for k in wc:
                cnt += wc[k]
            return cnt == 0

        def checkFrom(i: int, wc: dict) -> List[int]:
            ans = []
            for j in range(i, len(s) - wl + 1, wl):
                if s[j : j + wl] not in wc:
                    while i < j + wl:
                        if s[i : i + wl] in wc:
                            wc[s[i : i + wl]] += 1
                        i += wl
                    continue
                wc[s[j : j + wl]] -= 1
                while wc[s[j : j + wl]] < 0:
                    if s[i : i + wl] in wc:
                        wc[s[i : i + wl]] += 1
                    i += wl
                if allZero(wc):
                    ans.append(i)
            return ans

        idxs = []
        for i in range(wl):
            idxs += checkFrom(i, dict(wordDict))
        return idxs


# @lc code=end
